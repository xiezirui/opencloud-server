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

func GetUserByTelephone(telephone string, user *dao.User) bool {
	db.Where("telephone = ?", telephone).First(user)
	if user.ID != 0 {
		return true
	}
	return false
}
func AddNewUser(user *dao.User) {
	db.Create(user)
}
func GetUserById(id int, user *dao.User) {
	db.First(&user, id)
}
func UpDataNameById(id int, name string, userDto *dto.UserDto) {
	var user dao.User
	user.ID = uint(id)
	db.Model(&user).Update("name", name)
	dto.ToUserDto(user, userDto)
}
func UpDataPasswordById(id int, password string, userDto *dto.UserDto) {
	var user dao.User
	user.ID = uint(id)
	db.Model(&user).Update("password", password)
	dto.ToUserDto(user, userDto)
}
