/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   User     : ICI
   Name     : Ibrahim ÇOBANİ
   Date     : 22.07.2022
   Time     : 13:01
*/

package config

import "gin-gonic-web-Framework/models"

var Users []models.User

func InitMockData() {
	Users = []models.User{
		{
			ID:       1,
			Name:     "Ibrahim",
			UserName: "icobani",
			Password: "p@ssword",
		},
		{
			ID:       2,
			Name:     "Kubra",
			UserName: "kubra",
			Password: "p@ssword3",
		}}
}
