/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   User     : ICI
   Name     : Ibrahim ÇOBANİ
   Date     : 22.07.2022
   Time     : 13:43
*/

package rooter

import (
	"gin-gonic-web-Framework/controller"
	"github.com/gin-gonic/gin"
)

func SurveyRoot(api *gin.RouterGroup) {
	api.GET("/surveys", controller.GETUsers)
	api.GET("surveys/:id", controller.GETUserByID)
	api.POST("/surveys", controller.POSTUser)
	api.PUT("/surveys/:id", controller.PUTUser)
	api.DELETE("/surveys/:id", controller.DELETEUser)
}
