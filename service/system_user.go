// Package system_user @Author:冯铁城 [17615007230@163.com] 2025-07-30 19:41:24
package system_user

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net-project-edu_manage/dao"
	"net-project-edu_manage/dao/model/system"
	"net-project-edu_manage/dao/query"
	"net-project-edu_manage/model/dto"
	"net-project-edu_manage/util"
)

// Add 新增系统用户
func Add(c *gin.Context, systemUserDTO *dto.SystemUserDTO) error {
	return dao.Q.Transaction(func(tx *query.Query) error {

		//1.密码加密
		if password, err := util.HashPassword(systemUserDTO.Password); err != nil {
			return err
		} else {
			systemUserDTO.Password = password
		}

		//2.dto to po
		var systemUser system.SystemUser
		if err := copier.Copy(&systemUser, &systemUserDTO); err != nil {
			return err
		}

		//3.入库保存
		if err := tx.SystemUser.WithContext(c).Create(&systemUser); err != nil {
			return err
		}

		//4.ID回写,密码清空
		systemUserDTO.ID = systemUser.ID
		systemUserDTO.Password = ""

		//5.默认返回
		return nil
	})
}
