package dto

import (
	"opencloud-server/model/dao"
)

type UserDto struct {
	ID        uint
	Name      string
	Telephone string
}

func ToUserDto(user dao.User, userDto *UserDto) {
	userDto.ID = user.ID
	userDto.Name = user.Name
	userDto.Telephone = user.Telephone
}
