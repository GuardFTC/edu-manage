// Package dto @Author:冯铁城 [17615007230@163.com] 2025-08-04 16:26:37
package dto

import (
	"fmt"
	"time"
)

// AcademicYearDto 学年DTO
type AcademicYearDto struct {
	Name         string    `gorm:"column:name" json:"name" binding:"omitempty,max=32"`                         // 学年名称，例如 2025-2026
	StartDateStr string    `gorm:"column:start_date" json:"startDate" binding:"required,datetime=2006-01-02" ` // 学年开始日期
	EndDateStr   string    `gorm:"column:end_date" json:"endDate" binding:"required,datetime=2006-01-02"`      // 学年结束日期
	StartDate    time.Time `json:"-"`
	EndDate      time.Time `json:"-"`
	BaseDto
}

// ParseDate 解析时间
func (dto *AcademicYearDto) ParseDate() error {

	//1.解析开始时间
	startDate, err := time.Parse(time.DateOnly, dto.StartDateStr)
	if err != nil {
		return err
	}
	dto.StartDate = startDate

	//2.解析结束时间
	endDate, err := time.Parse(time.DateOnly, dto.EndDateStr)
	if err != nil {
		return err
	}
	dto.EndDate = endDate

	//3.默认返回
	return err
}

// FormatDate 格式化时间
func (dto *AcademicYearDto) FormatDate() {

	//1.解析开始时间
	dto.StartDateStr = dto.StartDate.Format(time.DateOnly)

	//2.解析结束时间
	dto.EndDateStr = dto.EndDate.Format(time.DateOnly)
}

// SetName 设置名称
func (dto *AcademicYearDto) SetName() {

	//1.如果名称不为空，直接返回
	if dto.Name != "" {
		return
	}

	//2.获取起止时间年份
	startYear := dto.StartDate.Year()
	endYear := dto.EndDate.Year()

	//3.拼接名称
	dto.Name = fmt.Sprintf("%d-%d学年", startYear, endYear)
}
