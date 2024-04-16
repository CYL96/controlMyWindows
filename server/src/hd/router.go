/**************************************************
*文件名：router.go
*内容简述：*
*文件历史：
author CYL96 创建 2024/3/26
**************************************************/

package hd

import (
	"os"

	"github.com/gin-gonic/gin"
)

func GinRouter(e *gin.Engine) {
	api := e.Group("/api")
	api.POST("/Exit", func(c *gin.Context) {
		os.Exit(0)
	})

	api.POST("/AddControlClass", AddControlClass)
	api.POST("/UpdateControlClass", UpdateControlClass)
	api.POST("/DeleteControlClass", DeleteControlClass)
	api.POST("/UpdateControlClassOrder", UpdateControlClassOrder)
	api.POST("/GetControlClassList", GetControlClassList)
	api.POST("/GetControlClassInfo", GetControlClassInfo)

	api.POST("/AddControlDetail", AddControlDetail)
	api.POST("/UpdateControlDetail", UpdateControlDetail)
	api.POST("/DeleteControlDetail", DeleteControlDetail)
	api.POST("/UpdateControlDetailOrder", UpdateControlDetailOrder)
	api.POST("/GetControlDetailList", GetControlDetailList)
	api.POST("/ExecControlDetail", ExecControlDetail)
	api.POST("/StopControlDetail", StopControlDetail)
	api.POST("/GetNowMousePosition", GetNowMousePosition)

	api.POST("/GetSystemConfig", GetSystemConfig)
	api.POST("/SetSystemConfig", SetSystemConfig)

	api.POST("/GetNowIconList", GetNowIconList)

}
