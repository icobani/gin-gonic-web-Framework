/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   User     : ICI
   Name     : Ibrahim ÇOBANİ
   Date     : 22.07.2022
   Time     : 13:36
*/

package rooter

import (
	"gin-gonic-web-Framework/controller"
	"github.com/gin-gonic/gin"
)

func UsersRoot(api *gin.RouterGroup) {
	api.GET("/users", controller.GETUsers)
	api.GET("users/:id", controller.GETUserByID)
	api.POST("/users", controller.POSTUser)
	api.PUT("/users/:id", controller.PUTUser)
	api.DELETE("/users/:id", controller.DELETEUser)
}
