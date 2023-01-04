package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"opencloud-server/database"
	"opencloud-server/model"
	"opencloud-server/response"
	"opencloud-server/utils"
)

func Register(context *gin.Context) {
	var user model.User

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

	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}

	database.AddNewUser(&newUser)
	response.Success(context, nil, "注册成功")
}
