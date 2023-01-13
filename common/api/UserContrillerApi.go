package api

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
)

func LoginFunction(telephone string, password string, context *gin.Context, user dao.User, userDto *dto.UserDto) {
	if len(telephone) != 11 {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	database.GetUserByTelephone(telephone, &user)
	if user.ID == 0 {
		response.Response(context, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}
	//密码解密
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
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

	dto.ToUserDto(user, userDto)

	response.Success(context, gin.H{"token": token, "user": userDto}, "登录成功")
	log.Printf("%v", userDto)
}
