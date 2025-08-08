// Package grade @Author:冯铁城 [17615007230@163.com] 2025-08-04 16:37:12
package grade

import (
	"net-project-edu_manage/internal/infrastructure/db"
	"net-project-edu_manage/internal/infrastructure/db/master/model"
	"net-project-edu_manage/internal/infrastructure/db/master/query"
	"net-project-edu_manage/internal/model/base"
	dtoPack "net-project-edu_manage/internal/model/dto/grade"
	reqPack "net-project-edu_manage/internal/model/request/grade"
	"net-project-edu_manage/internal/model/res"
	voPack "net-project-edu_manage/internal/model/vo/grade"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

// AcademicYearService 学年服务
type AcademicYearService struct {
	sync sync.Mutex //预留锁 并发高时使用
}

// Add 新增
func (s *AcademicYearService) Add(c *gin.Context, academicYearDto *dtoPack.AcademicYearDto) error {
	return db.GetDefaultQuery().Transaction(func(tx *query.Query) error {

		//1.设置名称
		academicYearDto.SetName()

		//2.设置创建人与修改人
		academicYearDto.SetCreateByAndUpdateBy(c)

		//3.dto to po
		var academicYear model.AcademicYear
		if err := copier.Copy(&academicYear, &academicYearDto); err != nil {
			return err
		}

		//4.保存
		if err := tx.AcademicYear.WithContext(c).Create(&academicYear); err != nil {
			return err
		}

		//5.ID回写
		academicYearDto.ID = academicYear.ID

		//6.默认返回
		return nil
	})
}

// Delete 删除学年
func (s *AcademicYearService) Delete(c *gin.Context, ids []string) error {
	return db.GetDefaultQuery().Transaction(func(tx *query.Query) error {

		//1.id string 转 int64
		intIds := cast.ToInt64Slice(ids)

		//2.删除学年
		if delRes, err := tx.AcademicYear.WithContext(c).Where(tx.AcademicYear.ID.In(intIds...)).Delete(); err != nil {
			return err
		} else {
			log.Printf("删除学年成功,删除数量:%d", delRes.RowsAffected)
		}

		//3.删除年级-学年关联
		if delRes, err := tx.GradeYear.WithContext(c).Where(
			tx.GradeYear.AcademicYearID.In(intIds...),
		).Delete(); err != nil {
			return err
		} else {
			log.Printf("删除年级-学年关联成功,删除数量:%d", delRes.RowsAffected)
			return nil
		}
	})
}

// Get 获取学年
func (s *AcademicYearService) Get(c *gin.Context, id string) (*dtoPack.AcademicYearDto, error) {

	//1.id string 转 int64
	intId := cast.ToInt64(id)

	//2.查询学年
	a := db.GetDefaultQuery().AcademicYear
	academicYear, err := a.WithContext(c).Where(a.ID.Eq(intId)).First()
	if err != nil {
		return nil, err
	}

	//3.po to dto
	var academicYearDTO dtoPack.AcademicYearDto
	if err = copier.Copy(&academicYearDTO, &academicYear); err != nil {
		return nil, err
	}

	//4.格式化时间
	academicYearDTO.FormatDate()

	//5.返回dto
	return &academicYearDTO, nil
}

// Update 修改学年
func (s *AcademicYearService) Update(c *gin.Context, id string, academicYearDto *dtoPack.AcademicYearDto) error {
	return db.GetDefaultQuery().Transaction(func(tx *query.Query) error {

		//1.id string 转 int64
		intId := cast.ToInt64(id)

		//2.查询学年
		academicYear, err := tx.AcademicYear.WithContext(c).Where(tx.AcademicYear.ID.Eq(intId)).First()
		if err != nil {
			return err
		}

		//3.设置ID
		academicYearDto.ID = intId

		//4.设置名称
		academicYearDto.SetName()

		//5.设置修改人
		academicYearDto.SetUpdateBy(c)

		//6.dto to po
		if err = copier.Copy(&academicYear, &academicYearDto); err != nil {
			return err
		}

		//7.更新
		if updateRes, err := tx.AcademicYear.WithContext(c).Where(tx.AcademicYear.ID.Eq(intId)).Updates(&academicYear); err != nil {
			return err
		} else {
			log.Printf("更新学年成功,更新数量:%d", updateRes.RowsAffected)
		}

		//8.返回
		return nil
	})
}

// Page 分页查询学年
func (s *AcademicYearService) Page(c *gin.Context, request *reqPack.AcademicYearRequest) (*res.PageResult[*voPack.AcademicYearVo], error) {

	//1.分页参数默认处理
	request.DefaultPage()

	//2.设置别名，利于后续Join查询
	ay := db.GetDefaultQuery().AcademicYear.As("ay")
	s1 := db.GetDefaultQuery().SystemUser.As("s1")
	s2 := db.GetDefaultQuery().SystemUser.As("s2")

	//3.封装查询参数
	context := ay.WithContext(c)
	if request.Name != "" {
		context = context.Where(ay.Name.Like("%" + request.Name + "%"))
	}
	if request.StartDateBegin != "" {
		location := cast.ToTimeInDefaultLocation(request.StartDateBegin, base.DefaultLoc)
		context = context.Where(ay.StartDate.Gte(location))
	}
	if request.StartDateEnd != "" {
		location := cast.ToTimeInDefaultLocation(request.StartDateEnd, base.DefaultLoc)
		context = context.Where(ay.StartDate.Lte(location))
	}
	if request.EndDateBegin != "" {
		location := cast.ToTimeInDefaultLocation(request.EndDateBegin, base.DefaultLoc)
		context = context.Where(ay.EndDate.Gte(location))
	}
	if request.EndDateEnd != "" {
		location := cast.ToTimeInDefaultLocation(request.EndDateEnd, base.DefaultLoc)
		context = context.Where(ay.EndDate.Lte(location))
	}

	//4.暂存总数查询参数
	countContext := context

	//5.设置查询字段，排序，分页参数
	context = context.
		Select(ay.ID, ay.Name, ay.StartDate, ay.EndDate, ay.CreatedAt, ay.UpdatedAt, s1.Name.As("created_user"), s2.Name.As("updated_user")).
		Join(s1, ay.CreatedBy.EqCol(s1.Email)).
		Join(s2, ay.UpdatedBy.EqCol(s2.Email)).
		Order(ay.ID.Desc()).
		Offset(request.GetSkip()).Limit(request.PageSize)

	//6.查询数据
	var academicYearVos []*voPack.AcademicYearVo
	if err := context.Scan(&academicYearVos); err != nil {
		return nil, err
	}

	//7.循环设置时间格式化模版
	for _, academicYearVo := range academicYearVos {
		academicYearVo.StartDate.Layout = time.DateOnly
		academicYearVo.EndDate.Layout = time.DateOnly
	}

	//8.查询总数
	total, err := countContext.Count()
	if err != nil {
		return nil, err
	}

	//9.封装分页结果
	return res.CreatePageResult[*voPack.AcademicYearVo](&request.Request, total, academicYearVos), nil
}

// Grades 获取指定学年的班级列表
func (s *AcademicYearService) Grades(c *gin.Context, id string) ([]*voPack.SimpleGradeVo, error) {

	//1.id string 转 int64
	intId := cast.ToInt64(id)

	//2.设置别名，利于后续Join查询
	g := db.GetDefaultQuery().Grade.As("g")
	gy := db.GetDefaultQuery().GradeYear.As("gy")

	//3.查询
	var gradeVos []*voPack.SimpleGradeVo
	err := g.WithContext(c).
		Select(g.ID, g.Name).
		Join(gy, g.ID.EqCol(gy.GradeID)).
		Where(gy.AcademicYearID.Eq(intId)).
		Order(g.ID.Desc()).
		Scan(&gradeVos)
	if err != nil {
		return nil, err
	}

	//4.空值处理
	if gradeVos == nil {
		return make([]*voPack.SimpleGradeVo, 0), nil
	}

	//5.返回
	return gradeVos, nil
}

// AddGrades 添加学年-年级关联
func (s *AcademicYearService) AddGrades(c *gin.Context, academicYearId string, dto *dtoPack.YearGradeDto) error {
	return db.GetDefaultQuery().Transaction(func(tx *query.Query) error {

		//1.id string 转 int64
		intAcademicYearId := cast.ToInt64(academicYearId)

		//2.确认学年存在
		if _, err := tx.AcademicYear.WithContext(c).Where(tx.AcademicYear.ID.Eq(intAcademicYearId)).First(); err != nil {
			return err
		}

		//3.筛选出真实存在的年级ID
		var existingGradeIds []int64
		if err := tx.Grade.WithContext(c).Where(tx.Grade.ID.In(dto.GradeIDs...)).Pluck(tx.Grade.ID, &existingGradeIds); err != nil {
			return err
		}

		//4.空值处理
		if existingGradeIds == nil {
			existingGradeIds = []int64{}
		}

		//5.重置年级ID集合
		dto.GradeIDs = existingGradeIds

		//6.如果没有任何一个年级存在，直接返回
		if len(existingGradeIds) == 0 {
			return nil
		}

		//7.设置创建人以及修改人
		dto.SetCreateByAndUpdateBy(c)

		//8.删除旧的关联关系
		if _, err := tx.GradeYear.WithContext(c).Where(tx.GradeYear.AcademicYearID.Eq(intAcademicYearId)).Delete(); err != nil {
			return err
		}

		//9.创建新的关联关系
		gradeYears := make([]*model.GradeYear, len(existingGradeIds))
		for i, gradeId := range existingGradeIds {
			gradeYears[i] = &model.GradeYear{
				AcademicYearID: intAcademicYearId,
				GradeID:        gradeId,
				CreatedBy:      dto.CreatedBy,
				UpdatedBy:      dto.UpdatedBy,
			}
		}

		//10.批量插入
		if err := tx.GradeYear.WithContext(c).Create(gradeYears...); err != nil {
			return err
		}

		//11.返回
		return nil
	})
}

// DeleteGrades 删除学年-年级关联
func (s *AcademicYearService) DeleteGrades(c *gin.Context, academicYearId string, gradeIds []string) error {
	return db.GetDefaultQuery().Transaction(func(tx *query.Query) error {

		//1.id string 转 int64
		intAcademicYearId := cast.ToInt64(academicYearId)
		intGradeIds := cast.ToInt64Slice(gradeIds)

		//2.删除关联
		if delRes, err := tx.GradeYear.WithContext(c).
			Where(tx.GradeYear.AcademicYearID.Eq(intAcademicYearId), tx.GradeYear.GradeID.In(intGradeIds...)).
			Delete(); err != nil {
			return err
		} else {
			log.Printf("删除年级-学年关联成功,删除数量:%d", delRes.RowsAffected)
		}

		//3.返回
		return nil
	})
}
