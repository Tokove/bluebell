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

// 升级版：可按照时间或分数来获取帖子
// 根据前端传来的参数动态获取列表
// 1.获取参宿
// 2.redis查询id列表
// 3.去数据库查询对应信息
func GetPostHandler2(c *gin.Context) {
	p := &model.ParamPostList{
		Page:  model.DefaultPage,
		Size:  model.DefaultSize,
		Order: model.OrderTime,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostHandler2 with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := service.GetPostListNew(p)
	if err != nil {
		zap.L().Error("service.GetPostList2 failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}


