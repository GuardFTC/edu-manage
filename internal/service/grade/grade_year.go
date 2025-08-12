// Package grade @Author:冯铁城 [17615007230@163.com] 2025-08-12 10:36:36
package grade

import (
	"errors"
	"net-project-edu_manage/internal/infrastructure/db"
	"net-project-edu_manage/internal/infrastructure/db/master/model"
	"net-project-edu_manage/internal/infrastructure/db/master/query"
	dtoPack "net-project-edu_manage/internal/model/dto/grade"
	reqPack "net-project-edu_manage/internal/model/request/grade"
	"net-project-edu_manage/internal/model/res"
	voPack "net-project-edu_manage/internal/model/vo/grade"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

// GradeYearService 学年服务
type GradeYearService struct {
	sync sync.Mutex //预留锁 并发高时使用
}

// Add 添加学年-年级关联
func (s *GradeYearService) Add(c *gin.Context, gradeYearDto *dtoPack.GradeYearDto) error {
	return db.GetDefaultQuery().Transaction(func(tx *query.Query) error {

		//1.校验DTO
		err := checkGradeYearDto(tx, c, gradeYearDto)
		if err != nil {
			return err
		}

		//2.设置创建人以及修改人
		gradeYearDto.SetCreateByAndUpdateBy(c)

		//3.创建新的关联关系
		gradeYear := &model.GradeYear{
			AcademicYearID: gradeYearDto.AcademicYearId,
			GradeID:        gradeYearDto.GradeId,
			CreatedBy:      gradeYearDto.CreatedBy,
			UpdatedBy:      gradeYearDto.UpdatedBy,
		}

		//4.保存
		if err = tx.GradeYear.WithContext(c).Create(gradeYear); err != nil {
			return err
		}

		//5.ID回写
		gradeYearDto.ID = gradeYear.ID

		//6.返回
		return nil
	})
}

// Delete 删除学年-年级关联
func (s *GradeYearService) Delete(c *gin.Context, ids []string) error {
	return db.GetDefaultQuery().Transaction(func(tx *query.Query) error {

		//1.id to int
		intIds := cast.ToInt64Slice(ids)

		//2.查询是否关联班级，如果是，无法删除
		count, err := tx.Class.WithContext(c).Where(tx.Class.GradeYearID.In(intIds...)).Count()
		if err != nil {
			return err
		}
		if count > 0 {
			return errors.New("grade and academic year associated classes! " + res.UnProcessTag)
		}

		//3.删除关联
		if delRes, err := tx.GradeYear.WithContext(c).Where(tx.GradeYear.ID.In(intIds...)).Delete(); err != nil {
			return err
		} else {
			log.Printf("删除年级-学年关联成功,删除数量:%d", delRes.RowsAffected)
			return nil
		}
	})
}

// Update 修改学年-年级关联
func (s *GradeYearService) Update(c *gin.Context, id string, gradeYearDto *dtoPack.GradeYearDto) error {
	return db.GetDefaultQuery().Transaction(func(tx *query.Query) error {

		//1.id to int
		intId := cast.ToInt64(id)

		//2.设置ID
		gradeYearDto.ID = intId

		//3.校验DTO
		err := checkGradeYearDto(tx, c, gradeYearDto)
		if err != nil {
			return err
		}

		//4.查询学年-年级关联
		gradeYear, err := tx.GradeYear.WithContext(c).Where(tx.GradeYear.ID.Eq(intId)).First()
		if err != nil {
			return err
		}

		//5.设置修改人
		gradeYearDto.SetUpdateBy(c)

		//6.dto to po
		if err = copier.Copy(&gradeYear, &gradeYearDto); err != nil {
			return err
		}

		//7.更新
		if updateRes, err := tx.GradeYear.WithContext(c).Where(tx.GradeYear.ID.Eq(intId)).Updates(&gradeYear); err != nil {
			return err
		} else {
			log.Printf("更新年级-学年关联成功,更新数量:%d", updateRes.RowsAffected)
		}

		//8.返回
		return nil
	})
}

// Page 分页查询学年-年级关联
func (s *GradeYearService) Page(c *gin.Context, request *reqPack.GradeYearRequest) (*res.PageResult[*voPack.GradeYearVo], error) {

	//1.分页参数默认处理
	request.DefaultPage()

	//2.设置别名，后续查询
	gy := db.GetDefaultQuery().GradeYear.As("gy")

	//3.封装查询参数
	context := gy.WithContext(c)
	if request.GradeId != 0 {
		context = context.Where(gy.GradeID.Eq(request.GradeId)).Order(gy.AcademicYearID.Desc())
	}
	if request.AcademicYearId != 0 {
		context = context.Where(gy.AcademicYearID.Eq(request.AcademicYearId)).Order(gy.GradeID.Desc())
	}

	//4.暂存总数查询参数
	countContext := context

	//5.设置查询字段，分页参数
	context = context.
		Select(gy.ID, gy.GradeID, gy.AcademicYearID).
		Offset(request.GetSkip()).Limit(request.PageSize)

	//6.查询数据
	var gradeYearVos []*voPack.GradeYearVo
	if err := context.Scan(&gradeYearVos); err != nil {
		return nil, err
	}

	//7.查询总数
	total, err := countContext.Count()
	if err != nil {
		return nil, err
	}

	//8.封装分页结果
	return res.CreatePageResult[*voPack.GradeYearVo](&request.Request, total, gradeYearVos), nil
}

// checkGradeYearDto 校验年级-学年关联DTO
func checkGradeYearDto(tx *query.Query, c *gin.Context, gradeYearDto *dtoPack.GradeYearDto) error {

	//1.确认学年存在
	if _, err := tx.AcademicYear.WithContext(c).Where(tx.AcademicYear.ID.Eq(gradeYearDto.AcademicYearId)).First(); err != nil {
		return err
	}

	//2.确认年级存在
	if _, err := tx.Grade.WithContext(c).Where(tx.Grade.ID.Eq(gradeYearDto.GradeId)).First(); err != nil {
		return err
	}

	//3.设置数量变量
	var count int64
	var err error

	//4.查询是否已进行关联
	if gradeYearDto.ID > 0 {
		count, err = tx.GradeYear.WithContext(c).Where(
			tx.GradeYear.ID.Neq(gradeYearDto.ID),
			tx.GradeYear.AcademicYearID.Eq(gradeYearDto.AcademicYearId),
			tx.GradeYear.GradeID.Eq(gradeYearDto.GradeId),
		).Count()
	} else {
		count, err = tx.GradeYear.WithContext(c).Where(
			tx.GradeYear.AcademicYearID.Eq(gradeYearDto.AcademicYearId),
			tx.GradeYear.GradeID.Eq(gradeYearDto.GradeId),
		).Count()
	}

	//5.异常以及数量判定
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("grade and academic year is associated! " + res.UnProcessTag)
	}

	//6.默认返回
	return err
}
