package route

import (
	"github.com/gin-gonic/gin"
	"github.com/wqh66886/past-present-future/controller"
	"github.com/wqh66886/past-present-future/mapper"
	"github.com/wqh66886/past-present-future/repository"
	"golang.org/x/net/context"
)

func NewRefreshTokenRouter(ctx context.Context, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(ctx)
	rt := controller.RefreshTokenController{
		RefreshTokenMapper: mapper.NewRefreshTokenMapper(ur),
	}
	group.POST("/refreshToken", rt.RefreshToken)
}
