package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wqh66886/past-present-future/common"
	"github.com/wqh66886/past-present-future/define"
	"github.com/wqh66886/past-present-future/domain"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type LoginController struct {
	LoginMapper domain.LoginUsecase
}

func (lc *LoginController) Login(c *gin.Context) {
	var request domain.LoginRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			Code:    http.StatusBadRequest,
			Message: "Invalid request",
			Data:    nil,
		})
		return
	}

	user, err := lc.LoginMapper.GetUserByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, common.Response{
			Code:    http.StatusNotFound,
			Message: "User not found",
			Data:    nil,
		})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		c.JSON(http.StatusUnauthorized, common.Response{
			Code:    http.StatusUnauthorized,
			Message: "Invalid password",
			Data:    nil,
		})
		return
	}

	accessToken, err := lc.LoginMapper.CreateAccessToken(user, define.Cfg.Auth.ExpireTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create access token",
			Data:    nil,
		})
		return
	}

	refreshToken, err := lc.LoginMapper.CreateRefreshToken(user, define.Cfg.Auth.RefreshTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create refresh token",
			Data:    nil,
		})
		return
	}

	loginResponse := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, loginResponse)
}
