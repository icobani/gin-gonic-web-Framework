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
	var returnValue models.User
	if userName, isExist := c.GetQuery("userName"); isExist {
		fmt.Println("userName:", userName)

		for _, user := range config.Users {
			if user.Name == userName {
				returnValue = user
			}
		}

		c.JSON(http.StatusOK, returnValue)
		return
	}

	c.JSON(http.StatusOK, config.Users)
}

func GETUserByID(c *gin.Context) {
	var returnValue models.User
	id_ := c.Param("id")
	id, err := strconv.ParseUint(id_, 10, 64)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, "invalid id")
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
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	config.Users = append(config.Users, user)
	c.JSON(http.StatusCreated, user)
}

func PUTUser(c *gin.Context) {
	var user models.User
	id_ := c.Param("id")
	id, err := strconv.ParseUint(id_, 10, 64)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
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
