package router

import (
	"bluebell_backend/controller"
	_ "bluebell_backend/docs" // 千万不要忘了导入把你上一步生成的docs
	"bluebell_backend/logger"
	"bluebell_backend/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	// ---------------------------------------------------------------------------------
	v1 := r.Group("/api/v1")
	v1.POST("/register", controller.RegisterHandler)
	v1.POST("/login", controller.LoginHandler)

	v1.Use(middleware.JWTAuthMiddleware())
	{
		v1.GET("/community", controller.GetCommunityHandler)
		v1.GET("/community/:id", controller.GetCommunityDetailHandler)

		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		v1.GET("/posts", controller.GetPostListHandler)
		v1.GET("/posts2", controller.GetPostListHandler2)

		v1.POST("/vote", controller.PostVoteHandler)
	}
	// ----------------------------------------------------------------------------------
	r.NoRoute(func(c *gin.Context) {
		controller.ResponseErrorWithMsg(c, controller.CodeNotFound, "404")
	})
	return r
}
