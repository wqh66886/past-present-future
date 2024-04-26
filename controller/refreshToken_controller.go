package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wqh66886/past-present-future/common"
	"github.com/wqh66886/past-present-future/define"
	"github.com/wqh66886/past-present-future/domain"
	"net/http"
)

type RefreshTokenController struct {
	RefreshTokenMapper domain.RefreshTokenUsecase
}

func (rtc *RefreshTokenController) RefreshToken(c *gin.Context) {
	var request domain.RefreshTokenRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.Response{Message: err.Error()})
		return
	}

	id, err := rtc.RefreshTokenMapper.ExtractIDFromToken(request.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, common.Response{Message: "User not found"})
		return
	}

	user, err := rtc.RefreshTokenMapper.GetUserByID(c, id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, common.Response{Message: "User not found"})
		return
	}

	accessToken, err := rtc.RefreshTokenMapper.CreateAccessToken(user, define.Cfg.Auth.ExpireTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
		return
	}

	refreshToken, err := rtc.RefreshTokenMapper.CreateRefreshToken(user, define.Cfg.Auth.RefreshTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
		return
	}

	refreshTokenResponse := domain.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, refreshTokenResponse)
}
