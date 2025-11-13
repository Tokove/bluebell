package controller

import (
	"bluebell_backend/dao/mysql"
	"bluebell_backend/service"
	"errors"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetCommunityHandler(c *gin.Context) {
	communityList, err := service.GetCommunityList()
	if err != nil {
		zap.L().Error("service.GetCommunityList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不要轻易将服务端错误暴漏
		return
	}
	ResponseSuccess(c, communityList)
}

func GetCommunityDetailHandler(c *gin.Context) {
	idStr := strings.TrimSpace(c.Param("id"))
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		zap.L().Error("get community detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	detail, err := service.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("get community detail failed", zap.Uint64("id", id), zap.Error(err))
		if errors.Is(err, mysql.ErrorInvalidID) {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, detail)
}
