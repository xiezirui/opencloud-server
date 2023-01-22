package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"opencloud-server/common"
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
	log.Printf("%v", user)

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
	database.GetUserByTelephone(telephone, &user)
	if user.ID != "" {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "用户已注册")
		return
	}

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(context, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}

	newUser := dao.User{
		ID:        utils.GetUUID(),
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}

	log.Printf("%v", newUser)

	database.AddNewUser(&newUser)

	//自动登录
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(context, http.StatusInternalServerError, 500, nil, "token发放错误")
		log.Printf("token generate error : %v", err)
		return
	}

	dto.ToUserDto(newUser, &userDto)

	response.Success(context, gin.H{"token": token, "user": userDto}, "登录成功")
	log.Printf("%v", userDto)

}

func Login(context *gin.Context) {
	var user dao.User
	var dbUser dao.User
	var userDto dto.UserDto

	context.Bind(&user)

	if len(user.Telephone) != 11 {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	database.GetUserByTelephone(user.Telephone, &dbUser)
	if dbUser.ID == "" {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}
	//密码解密
	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "密码错误")
		return
	}
	//发放 token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(context, http.StatusInternalServerError, 500, nil, "token发放错误")
		log.Printf("token generate error : %v", err)
		return
	}

	dto.ToUserDto(dbUser, &userDto)

	response.Success(context, gin.H{"token": token, "user": userDto}, "登录成功")
	log.Printf("%v", dbUser)

}

func UpdataName(context *gin.Context) {
	var userDto dto.UserDto
	var upNameDto dto.UpNameDto
	context.Bind(&upNameDto)

	if len(upNameDto.Name) == 0 {
		response.Fail(context, nil, "昵称不能为空")
		return
	}
	database.UpDataNameById(upNameDto.ID, upNameDto.Name, &userDto)
	response.Success(context, gin.H{"user": userDto}, "修改成功")
}

func UpdataPassword(context *gin.Context) {
	var userDto dto.UserDto
	var upPasswordDto dto.UpPasswordDto
	context.Bind(&upPasswordDto)
	if upPasswordDto.Password == "" {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "密码不能为空")
		return
	}

	hasedPassword, _ := bcrypt.GenerateFromPassword([]byte(upPasswordDto.Password), bcrypt.DefaultCost)

	database.UpDataPasswordById(upPasswordDto.ID, string(hasedPassword), &userDto)

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

	database.GetUserById(OldPassowrdDto.ID, &user)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(OldPassowrdDto.OldPassword))

	if err != nil {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "密码错误")
		return
	}

	response.Success(context, nil, "密码验证成功")
}
