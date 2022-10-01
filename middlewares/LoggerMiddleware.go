/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   User     : ICI
   Name     : Ibrahim ÇOBANİ
   Date     : 01.10.2022
   Time     : 11:23
*/

package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

//Logger  Big Query Log
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request
		t := time.Now()
		log.Println("on before request")
		c.Next()
		log.Println("on after request")
		// after request
		latency := time.Since(t)
		log.Println("*** Latency :", latency.Milliseconds())
	}
}
