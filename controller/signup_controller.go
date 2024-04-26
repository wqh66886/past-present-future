package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wqh66886/past-present-future/common"
	"github.com/wqh66886/past-present-future/define"
	"github.com/wqh66886/past-present-future/domain"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type SignupController struct {
	SignUpMapper domain.SignupUsecase
}

func (sc *SignupController) Signup(c *gin.Context) {
	var request domain.SignUpRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.Response{Message: err.Error()})
		return
	}

	_, err = sc.SignUpMapper.GetUserByEmail(c, request.Email)
	if err == nil {
		c.JSON(http.StatusConflict, common.Response{Message: "User already exists with the given email"})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
		return
	}

	request.Password = string(encryptedPassword)

	user := domain.User{
		ID:         define.GetUUID(),
		Name:       request.Name,
		Email:      request.Email,
		Password:   request.Password,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	err = sc.SignUpMapper.Create(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
		return
	}

	accessToken, err := sc.SignUpMapper.CreateAccessToken(&user, define.Cfg.Auth.ExpireTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
		return
	}

	refreshToken, err := sc.SignUpMapper.CreateRefreshToken(&user, define.Cfg.Auth.RefreshTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.Response{Message: err.Error()})
		return
	}

	signupResponse := domain.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, signupResponse)
}
