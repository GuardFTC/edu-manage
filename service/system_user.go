// Package system_user @Author:冯铁城 [17615007230@163.com] 2025-07-30 19:41:24
package system_user

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"net-project-edu_manage/common/util"
	"net-project-edu_manage/core/db"
	"net-project-edu_manage/dao/model/system"
	"net-project-edu_manage/dao/query"
	"net-project-edu_manage/model/dto"
)

// Add 新增系统用户
func Add(c *gin.Context, systemUserDTO *dto.SystemUserDTO) error {
	return db.Q.Transaction(func(tx *query.Query) error {

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

// Delete 删除系统用户
func Delete(c *gin.Context, ids []string) error {
	return db.Q.Transaction(func(tx *query.Query) error {

		//1.id string 转 int64
		intIds := cast.ToInt64Slice(ids)

		//2.删除系统用户
		if delRes, err := tx.SystemUser.WithContext(c).Where(tx.SystemUser.ID.In(intIds...)).Delete(); err != nil {
			return err
		} else {
			log.Printf("删除系统用户成功,删除数量:%d", delRes.RowsAffected)
			return nil
		}
	})
}

// Get 获取系统用户
func Get(c *gin.Context, id string) (*dto.SystemUserDTO, error) {

	//1.id string 转 int64
	intId := cast.ToInt64(id)

	//2.查询系统用户
	systemUser, err := db.Q.SystemUser.WithContext(c).Where(db.Q.SystemUser.ID.Eq(intId)).First()
	if err != nil {
		return nil, err
	}

	//3.po to dto
	var systemUserDTO dto.SystemUserDTO
	if err = copier.Copy(&systemUserDTO, &systemUser); err != nil {
		return nil, err
	}

	//4.返回dto
	return &systemUserDTO, nil
}

// Update 修改系统用户
func Update(c *gin.Context, id string, systemUserDTO *dto.SystemUserDTO) error {
	return db.Q.Transaction(func(tx *query.Query) error {

		//1.id string 转 int64
		intId := cast.ToInt64(id)

		//2.查询系统用户
		systemUser, err := db.Q.SystemUser.WithContext(c).Where(db.Q.SystemUser.ID.Eq(intId)).First()
		if err != nil {
			return err
		}

		//3.dto to po
		if err = copier.Copy(&systemUser, &systemUserDTO); err != nil {
			return err
		}

		//4.更新
		if updateRes, err := tx.SystemUser.WithContext(c).Where(tx.SystemUser.ID.Eq(intId)).Updates(&systemUser); err != nil {
			return err
		} else {
			log.Printf("更新系统用户成功,更新数量:%d", updateRes.RowsAffected)
		}

		//5.ID回写
		systemUserDTO.ID = intId

		//6.返回
		return nil
	})
}
