package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/wqh66886/past-present-future/define"
	"github.com/wqh66886/past-present-future/errcode"
	"net/http"
	"strings"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if len(authHeader) == 0 {
			ctx.JSON(http.StatusUnauthorized, errcode.UnauthorizedTokenError)
			ctx.Abort()
		}
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := define.IsAuthorized(authToken)
			if authorized {
				userId, err := define.ExtractIDFromToken(authToken)
				if err != nil {
					ctx.JSON(http.StatusUnauthorized, errcode.NewError(401, err.Error()))
					ctx.Abort()
					return
				}
				ctx.Set("x-user-id", userId)
				ctx.Next()
				return
			}
			ctx.JSON(http.StatusUnauthorized, errcode.NewError(401, err.Error()))
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusUnauthorized, errcode.NewError(401, "Not authorized"))
		ctx.Abort()
	}
}
