// Package class @Author:冯铁城 [17615007230@163.com] 2025-08-11 16:29:00
package class

import (
	dtoPack "net-project-edu_manage/internal/model/dto/class"
	"net-project-edu_manage/internal/model/res"
	"net-project-edu_manage/internal/service/class"

	"github.com/gin-gonic/gin"
)

// classService 班级服务
var classService = new(class.ClassService)

// AddClass 添加班级
func AddClass(c *gin.Context) {

	//1.创建DTO
	var classDto dtoPack.ClassDto

	//2.校验参数并绑定
	if err := c.ShouldBindJSON(&classDto); err != nil {
		res.FailResToC(c, res.BadRequestFail, err.Error())
		return
	}

	//3.保存
	if err := classService.Add(c, &classDto); err != nil {
		res.FailResToCByMsg(c, err.Error())
		return
	}

	//4.返回
	res.SuccessResToC(c, res.CreateSuccess, classDto)

}

// DeleteClass 删除班级
func DeleteClass(c *gin.Context) {

	//1.获取参数
	ids := c.QueryArray("id")
	if len(ids) == 0 {
		res.FailResToC(c, res.BadRequestFail, "参数为空")
		return
	}

	//2.删除班级
	if err := classService.Delete(c, ids); err != nil {
		res.FailResToCByMsg(c, err.Error())
		return
	}

	//3.返回
	res.SuccessResToC(c, res.DeleteSuccess, nil)
}

//// GetClass 获取班级
//func GetClass(c *gin.Context) {
//
//	//1.获取路径参数班级ID
//	id := c.Param("id")
//
//	//2.查询班级
//	classDto, err := classService.Get(c, id)
//	if err != nil {
//		res.FailResToCByMsg(c, err.Error())
//	}
//
//	//3.返回
//	res.SuccessResToC(c, res.QuerySuccess, classDto)
//}
