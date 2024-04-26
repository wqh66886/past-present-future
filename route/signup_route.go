package route

import (
	"github.com/gin-gonic/gin"
	"github.com/wqh66886/past-present-future/controller"
	"github.com/wqh66886/past-present-future/mapper"
	"github.com/wqh66886/past-present-future/repository"
	"golang.org/x/net/context"
)

func NewSignupRouter(ctx context.Context, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(ctx)
	su := &controller.SignupController{
		SignUpMapper: mapper.NewSignupMapper(ur),
	}
	group.POST("/signup", su.Signup)
}
