package database

import (
	"gorm.io/gorm"
	"opencloud-server/model/dao"
	"opencloud-server/model/dto"
)

var db *gorm.DB

func init() {
	db = InitDatabase()
	SyncStruct(db)
}

func GetUserByTelephone(telephone string, user *dao.User) {
	db.Where("telephone = ?", telephone).First(user)
}
func AddNewUser(user *dao.User) {
	db.Create(user)
}
func GetUserById(id string, user *dao.User) {
	db.First(&user, "id = ?", id)
}
func UpDataNameById(id string, name string, userDto *dto.UserDto) {
	var user dao.User
	user.ID = id
	db.Model(&user).Update("name", name)
	dto.ToUserDto(user, userDto)
}
func UpDataPasswordById(id string, password string, userDto *dto.UserDto) {
	var user dao.User
	user.ID = id
	db.Model(&user).Update("password", password)
	dto.ToUserDto(user, userDto)
}
