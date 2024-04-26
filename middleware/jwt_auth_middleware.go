package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/wqh66886/past-present-future/common"
	"github.com/wqh66886/past-present-future/define"
	"net/http"
	"strings"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if len(authHeader) == 0 {
			ctx.JSON(http.StatusUnauthorized, common.Response{
				Code:    401,
				Message: "Not authorized",
			})
			ctx.Abort()
		}
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := define.IsAuthorized(authToken)
			if authorized {
				userId, err := define.ExtractIDFromToken(authToken)
				if err != nil {
					ctx.JSON(http.StatusUnauthorized, common.Response{
						Code:    401,
						Message: err.Error(),
					})
					ctx.Abort()
					return
				}
				ctx.Set("x-user-id", userId)
				ctx.Next()
				return
			}
			ctx.JSON(http.StatusUnauthorized, common.Response{
				Code:    401,
				Message: err.Error(),
			})
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusUnauthorized, common.Response{
			Code:    401,
			Message: "Not authorized",
		})
		ctx.Abort()
	}
}
