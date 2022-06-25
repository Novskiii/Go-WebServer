package controller

import (
	"awesomeProject/dao/mysql"
	"awesomeProject/logic"
	"awesomeProject/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func SignUpHandler(c *gin.Context) {
	//1.获取和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	if len(p.Username) < 6 || len(p.Password) < 8 {
		zap.L().Error("SignUp with invalid param")
		ResponseErrorWithMsg(c, CodeInvalidParam, "用户名长度应大于6位，密码长度应大于8位")
		return
	}
	fmt.Println(p)
	//2.业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
		return
	}
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.login failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, gin.H{
		"user_id":   strconv.FormatInt(user.UserID, 10),
		"user_name": user.Username,
		"token":     user.Token,
	})
}
