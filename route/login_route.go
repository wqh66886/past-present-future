package route

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/wqh66886/past-present-future/controller"
	"github.com/wqh66886/past-present-future/mapper"
	"github.com/wqh66886/past-present-future/repository"
)

func NewLoginRouter(ctx context.Context, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(ctx)
	lc := &controller.LoginController{
		LoginMapper: mapper.NewLoginMapper(ur),
	}
	group.POST("/login", lc.Login)
}
