package controller

import (
	"bluebell_backend/dao/mysql"
	"bluebell_backend/model"
	"bluebell_backend/service"
	"errors"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func CreatePostHandler(c *gin.Context) {
	p := new(model.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJSON(p) error", zap.Any("error", err))
		zap.L().Error("create post with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// 获取当前用户的ID
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNotLogin)
		return
	}
	p.AuthorID = userID

	if err := service.CreatePost(p); err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

func GetPostDetailHandler(c *gin.Context) {
	idStr := strings.TrimSpace(c.Param("id"))
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	detail, err := service.GetPostDetial(id)
	if err != nil {
		zap.L().Error("get post detail failed", zap.Uint64("id", id), zap.Error(err))
		if errors.Is(err, mysql.ErrorInvalidID) {
			ResponseError(c, CodeInvalidParam)
			return
		}
	}
	ResponseSuccess(c, detail)
}

func GetPostHandler(c *gin.Context) {
	page, size := getPageInfo(c)

	data, err := service.GetPostList(page, size)
	if err != nil {
		zap.L().Error("service.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}
	ResponseSuccess(c, data)
}
