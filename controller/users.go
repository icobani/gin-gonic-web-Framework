/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   User     : ICI
   Name     : Ibrahim ÇOBANİ
   Date     : 22.07.2022
   Time     : 13:00
*/

package controller

import (
	"fmt"
	"gin-gonic-web-Framework/config"
	"gin-gonic-web-Framework/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GETUsers(c *gin.Context) {
	c.JSON(http.StatusOK, config.Users)
}

func GETUserByID(c *gin.Context) {
	var returnValue models.User
	id_ := c.Param("id")
	id, err := strconv.ParseUint(id_, 10, 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(id, id_)
	for _, user := range config.Users {
		if user.ID == id {
			returnValue = user
		}
	}
	if returnValue.ID == 0 {
		c.JSON(http.StatusNotFound, "Kayıt bulunamadı.")
		return
	}
	c.JSON(http.StatusOK, returnValue)

}

func POSTUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	config.Users = append(config.Users, user)
	c.JSON(http.StatusOK, user)
}

func PUTUser(c *gin.Context) {
	var user models.User
	id_ := c.Param("id")
	id, err := strconv.ParseUint(id_, 10, 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.BindJSON(&user)
	for index, item := range config.Users {
		if item.ID == id {
			config.Users[index] = user
		}
	}
	c.JSON(http.StatusOK, user)
}

func DELETEUser(c *gin.Context) {
	var returnValue models.User
	id_ := c.Param("id")
	id, err := strconv.ParseUint(id_, 10, 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	for index, item := range config.Users {
		if item.ID == id {
			returnValue = item
			config.Users = append(config.Users[:index], config.Users[index+1:]...)
		}
	}
	if returnValue.ID == 0 {
		c.JSON(http.StatusNotFound, "Kayıt bulunamadı.")
		return
	}
	c.JSON(http.StatusOK, returnValue)
}
