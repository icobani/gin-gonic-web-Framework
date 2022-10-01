/*
	Standart CRUD Operasyonları ==> CREATE, READ, UPDATE, DELETE
	GET 	: Listleme için.
	GET 	: Kaydı getirmek için
	POST 	: Insert
	PUT		: Update
	DELETE	: DELETE


*/
package main

import (
	"fmt"
	"gin-gonic-web-Framework/config"
	"gin-gonic-web-Framework/middlewares"
	"gin-gonic-web-Framework/rooter"
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	config.InitMockData()
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.BestCompression))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.LoadHTMLGlob("templates/*")
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.POST("/upload", func(c *gin.Context) {
		// Single file
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// Upload the file to specific dst.
		dst := "./upload/" + file.Filename
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	r.Static("/upload", "./upload")

	r.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Minnak website",
			"name":  "Ibrahim",
		})
	})

	api := r.Group("")
	api.Use(middlewares.Logger())
	rooter.UsersRoot(api)
	rooter.SurveyRoot(api)
	//log.Fatal(autotls.Run(r, "example1.com", "example2.com"))
	err := r.Run(":8080")
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
