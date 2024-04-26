package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wqh66886/past-present-future/domain"
	"github.com/wqh66886/past-present-future/errcode"
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
		c.JSON(http.StatusBadRequest, errcode.NewError(400, "Invalid request body"))
		return
	}

	user, err := lc.LoginMapper.GetUserByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, errcode.NewError(400, "User not found with the given email"))
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		c.JSON(http.StatusUnauthorized, errcode.NewError(400, "Invalid credentials"))
		return
	}

	accessToken, err := lc.LoginMapper.CreateAccessToken(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errcode.NewError(400, "Failed to create access token"))
		return
	}

	refreshToken, err := lc.LoginMapper.CreateRefreshToken(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errcode.NewError(400, "Failed to create refresh token"))
		return
	}

	loginResponse := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, loginResponse)
}
