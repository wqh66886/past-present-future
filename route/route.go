package route

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/wqh66886/past-present-future/middleware"
)

func Setup(gin *gin.Engine) {
	publicRouter := gin.Group("")
	ctx := context.Background()
	NewSignupRouter(ctx, publicRouter)
	NewLoginRouter(ctx, publicRouter)
	NewRefreshTokenRouter(ctx, publicRouter)
	// 鉴权开启
	protectedRouter := gin.Group("")
	protectedRouter.Use(middleware.JwtAuthMiddleware())

}
