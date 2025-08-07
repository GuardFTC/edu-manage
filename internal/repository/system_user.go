// Package repository @Author:冯铁城 [17615007230@163.com] 2025-08-06 15:02:45
package repository

import (
	"net-project-edu_manage/internal/infrastructure/db"
	"net-project-edu_manage/internal/infrastructure/db/master/model"

	"github.com/gin-gonic/gin"
)

// SystemUserRepository 系统用户接口
type SystemUserRepository interface {
	GetByAccount(account string) (*model.SystemUser, error)
}

// SystemUserRepo 系统用户接口实现
type SystemUserRepo struct{}

// NewSystemUserRepository 创建系统用户接口实例
func NewSystemUserRepository() *SystemUserRepo {
	return &SystemUserRepo{}
}

// GetByAccount 根据账号获取系统用户
func (r *SystemUserRepo) GetByAccount(c *gin.Context, account string) (*model.SystemUser, error) {
	systemUser := db.GetDefaultQuery().SystemUser
	return systemUser.WithContext(c).
		Where(systemUser.Name.Eq(account)).
		Or(systemUser.Email.Eq(account)).
		First()
}
