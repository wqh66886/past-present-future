package route

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/wqh66886/past-present-future/middleware"
)

func Setup(gin *gin.Engine) {
	publicRouter := gin.Group("")
	ctx := context.Background()
	NewLoginRouter(ctx, publicRouter)

	// 鉴权开启
	gin.Use(middleware.JwtAuthMiddleware())
}
