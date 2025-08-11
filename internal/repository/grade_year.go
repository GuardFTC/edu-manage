// Package repository @Author:冯铁城 [17615007230@163.com] 2025-08-06 15:02:45
package repository

import (
	"net-project-edu_manage/internal/infrastructure/db"

	"github.com/gin-gonic/gin"
)

// GradeYearRepository 学年-年级接口
type GradeYearRepository interface {
	SelectClassCountByYearId(c *gin.Context, yearId int64) (int64, error)
	SelectClassCountByGradeId(c *gin.Context, yearId int64) (int64, error)
}

// GradeYearRepo 学年-年级接口实现
type GradeYearRepo struct{}

// NewGradeYearRepository 创建学年-年级接口实例
func NewGradeYearRepository() *GradeYearRepo {
	return &GradeYearRepo{}
}

// SelectClassCountByYearId 根据学年ID查询班级数量
func (r *GradeYearRepo) SelectClassCountByYearId(c *gin.Context, yearIds []int64) (int64, error) {

	//1.设置别名
	query := db.GetDefaultQuery()
	gy := query.GradeYear.As("gy")
	cl := query.Class.As("c")

	//2.查询
	return gy.WithContext(c).Join(cl, gy.ID.EqCol(cl.GradeYearID)).Where(gy.AcademicYearID.In(yearIds...)).Count()
}

// SelectClassCountByGradeId 根据年级ID查询班级数量
func (r *GradeYearRepo) SelectClassCountByGradeId(c *gin.Context, gradeIds []int64) (int64, error) {

	//1.设置别名
	query := db.GetDefaultQuery()
	gy := query.GradeYear.As("gy")
	cl := query.Class.As("c")

	//2.查询
	return gy.WithContext(c).Join(cl, gy.ID.EqCol(cl.GradeYearID)).Where(gy.GradeID.In(gradeIds...)).Count()
}
