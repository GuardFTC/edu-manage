// Package grade @Author:冯铁城 [17615007230@163.com] 2025-08-12 10:36:36
package grade

import (
	"errors"
	"net-project-edu_manage/internal/infrastructure/db"
	"net-project-edu_manage/internal/infrastructure/db/master/model"
	"net-project-edu_manage/internal/infrastructure/db/master/query"
	"net-project-edu_manage/internal/model/base"
	"net-project-edu_manage/internal/model/res"
	"net-project-edu_manage/internal/repository"
	"sync"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

// 年级-学年DB操作
var gradeYearRepo = repository.NewGradeYearRepository()

// GradeYearService 学年服务
type GradeYearService struct {
	sync sync.Mutex //预留锁 并发高时使用
}

// GetGradesByYearId 根据学年ID查询年级
func (s *GradeYearService) GetGradesByYearId(c *gin.Context, id string) ([]*base.SimpleVo, error) {

	//1.id string 转 int64
	intId := cast.ToInt64(id)

	//2.设置别名，利于后续Join查询
	g := db.GetDefaultQuery().Grade.As("g")
	gy := db.GetDefaultQuery().GradeYear.As("gy")

	//3.查询
	var simpleVos []*base.SimpleVo
	err := g.WithContext(c).
		Select(g.ID, g.Name).
		Join(gy, g.ID.EqCol(gy.GradeID)).
		Where(gy.AcademicYearID.Eq(intId)).
		Order(g.ID.Desc()).
		Scan(&simpleVos)
	if err != nil {
		return nil, err
	}

	//4.空值处理
	if simpleVos == nil {
		return make([]*base.SimpleVo, 0), nil
	}

	//5.返回
	return simpleVos, nil
}

// GetYearsByGradeId 根据年级ID查询学年
func (s *GradeYearService) GetYearsByGradeId(c *gin.Context, id string) ([]*base.SimpleVo, error) {

	//1.id string 转 int64
	intId := cast.ToInt64(id)

	//2.设置别名，利于后续Join查询
	ay := db.GetDefaultQuery().AcademicYear.As("ay")
	gy := db.GetDefaultQuery().GradeYear.As("gy")

	//3.查询
	var simpleVos []*base.SimpleVo
	err := ay.WithContext(c).
		Select(ay.ID, ay.Name).
		Join(gy, ay.ID.EqCol(gy.AcademicYearID)).
		Where(gy.GradeID.Eq(intId)).
		Order(ay.ID.Desc()).
		Scan(&simpleVos)
	if err != nil {
		return nil, err
	}

	//4.空值处理
	if simpleVos == nil {
		return make([]*base.SimpleVo, 0), nil
	}

	//5.返回
	return simpleVos, nil
}

// AddGradeYear 添加学年-年级关联
func (s *GradeYearService) AddGradeYear(c *gin.Context, academicYearId int64, gradeId int64) error {
	return db.GetDefaultQuery().Transaction(func(tx *query.Query) error {

		//1.确认学年存在
		if _, err := tx.AcademicYear.WithContext(c).Where(tx.AcademicYear.ID.Eq(academicYearId)).First(); err != nil {
			return err
		}

		//2.确认年级存在
		if _, err := tx.Grade.WithContext(c).Where(tx.Grade.ID.Eq(gradeId)).First(); err != nil {
			return err
		}

		//3.查询是否已进行关联
		count, err := tx.GradeYear.WithContext(c).Where(
			tx.GradeYear.AcademicYearID.Eq(academicYearId),
			tx.GradeYear.GradeID.Eq(gradeId),
		).Count()
		if err != nil {
			return err
		}
		if count > 0 {
			return errors.New("grade and academic year is associated! " + res.UnProcessTag)
		}

		//4.设置创建人以及修改人
		dto := new(base.Dto)
		dto.SetCreateByAndUpdateBy(c)

		//5.创建新的关联关系
		gradeYear := &model.GradeYear{
			AcademicYearID: academicYearId,
			GradeID:        gradeId,
			CreatedBy:      dto.CreatedBy,
			UpdatedBy:      dto.UpdatedBy,
		}

		//6.保存
		if err = tx.GradeYear.WithContext(c).Create(gradeYear); err != nil {
			return err
		}

		//7.返回
		return nil
	})
}

// DeleteGradeYear 删除学年-年级关联
func (s *GradeYearService) DeleteGradeYear(c *gin.Context, academicYearId int64, gradeId int64) error {
	return db.GetDefaultQuery().Transaction(func(tx *query.Query) error {

		//1.查询是否关联班级，如果是，无法删除
		count, err := gradeYearRepo.SelectClassCountByGradeYearId(c, academicYearId, gradeId)
		if err != nil {
			return err
		}
		if count > 0 {
			return errors.New("grade and academic year associated classes! " + res.UnProcessTag)
		}

		//2.删除关联
		if delRes, err := tx.GradeYear.WithContext(c).
			Where(tx.GradeYear.AcademicYearID.Eq(academicYearId), tx.GradeYear.GradeID.Eq(gradeId)).
			Delete(); err != nil {
			return err
		} else {
			log.Printf("删除年级-学年关联成功,删除数量:%d", delRes.RowsAffected)
		}

		//3.返回
		return nil
	})
}
