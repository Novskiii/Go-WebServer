package logic

import (
	"awesomeProject/dao/mysql"
	"awesomeProject/models"
	"awesomeProject/pkg/jwt"
	"awesomeProject/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) error {
	//判断用户存不存在
	err := mysql.CheckUserExist(p.Username)
	if err != nil {
		return err
	}
	//生成UID
	userID := snowflake.GetID()

	user := models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//保存进数据库
	return mysql.InsertUser(&user)
}

func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	err = mysql.Login(user)
	if err != nil {
		return nil, err
	}
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return nil, err
	}
	user.Token = token
	return
}
