/*
   B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
   User     : ICI
   Name     : Ibrahim ÇOBANİ
   Date     : 22.07.2022
   Time     : 12:36
*/

package models

type User struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
