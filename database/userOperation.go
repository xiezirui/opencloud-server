package database

import (
	"gorm.io/gorm"
	"opencloud-server/model"
)

var db *gorm.DB

func init() {
	db = InitDatabase()
	SyncStruct(db)
}

func GetUserByTelephone(telephone string, user *model.User) bool {
	db.Where("telephone = ?", telephone).First(user)
	if user.ID != 0 {
		return true
	}
	return false
}
func AddNewUser(user *model.User) {
	db.Create(user)
}
func GetUserById(id int, user *model.User) {
	db.First(&user, id)
}
