package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"opencloud-server/common/api"
	"opencloud-server/database"
	"opencloud-server/model/dao"
	"opencloud-server/model/dto"
	"opencloud-server/response"
	"opencloud-server/utils"
)

func Register(context *gin.Context) {
	var user dao.User
	var userDto dto.UserDto

	context.Bind(&user)

	name := user.Name
	telephone := user.Telephone
	password := user.Password

	if len(telephone) != 11 {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "密码不得小于6位")
		return
	}
	if len(name) == 0 {
		name = utils.GetRandomString(10)
	}

	if database.GetUserByTelephone(telephone, &user) {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "用户已注册")
		return
	}

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(context, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}

	newUser := dao.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}

	database.AddNewUser(&newUser)

	//自动登录
	api.LoginFunction(telephone, password, context, user, &userDto)

}

func Login(context *gin.Context) {
	var user dao.User
	var userDto dto.UserDto

	context.Bind(&user)

	telephone := context.PostForm("telephone")
	password := context.PostForm("password")

	api.LoginFunction(telephone, password, context, user, &userDto)

}

func UpdataName(context *gin.Context) {
	var userDto dto.UserDto
	var upNameDto dto.UpNameDto
	context.Bind(&upNameDto)

	if len(upNameDto.Name) == 0 {
		response.Fail(context, nil, "昵称不能为空")
		return
	}
	database.UpDataNameById(int(upNameDto.ID), upNameDto.Name, &userDto)
	response.Success(context, gin.H{"user": userDto}, "修改成功")
}

func UpdataPassword(context *gin.Context) {
	var userDto dto.UserDto
	var upPasswordDto dto.UpPasswordDto
	context.Bind(&upPasswordDto)
	if len(upPasswordDto.Password) == 0 {
		response.Fail(context, nil, "昵称不能为空")
		return
	}

	hasedPassword, _ := bcrypt.GenerateFromPassword([]byte(upPasswordDto.Password), bcrypt.DefaultCost)

	database.UpDataPasswordById(int(upPasswordDto.ID), string(hasedPassword), &userDto)

	response.Success(context, nil, "修改成功")
}

func CheckOldPassword(context *gin.Context) {
	var OldPassowrdDto dto.CheckOldPassowrdDto
	var user dao.User
	context.Bind(&OldPassowrdDto)

	if len(OldPassowrdDto.OldPassword) == 0 {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "密码不能为空")
		return
	}

	database.GetUserById(int(OldPassowrdDto.ID), &user)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(OldPassowrdDto.OldPassword))

	if err != nil {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "密码错误")
		return
	}

	response.Success(context, nil, "密码验证成功")
}
