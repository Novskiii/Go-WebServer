package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

const ContextUserIDKey = "userId"

var ErrorUserNotLogin = errors.New("用户未登陆")

//getCurrentUser 获取当前登陆用户ID
func getCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(ContextUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
	}
	return
}

func getPageInfo(c *gin.Context) (int64, int64) {
	pageStr := c.Query("offset")
	sizeStr := c.Query("limit")

	var (
		page int64
		size int64
		err  error
	)
	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	return page, size
}
