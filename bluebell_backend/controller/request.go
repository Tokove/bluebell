package controller

import (
	"bluebell_backend/dao/mysql"
	"strconv"

	"github.com/gin-gonic/gin"
)

const ContextUserIDKey = "userID"

func getCurrentUserID(c *gin.Context) (userID uint64, err error) {
	uid, ok := c.Get(ContextUserIDKey)
	if !ok {
		err = mysql.ErrorNotLogin
		return
	}
	userID, ok = uid.(uint64)
	if !ok {
		err = mysql.ErrorNotLogin
		return
	}
	return userID, nil
}

func getPageInfo(c *gin.Context) (int64, int64) {
	pageStr := c.Query("page")
	sizeStr := c.Query("size")
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
