package define

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	model2 "github.com/wqh66886/past-present-future/model"
	"gorm.io/gorm"
	"time"
)

var (
	DB  *gorm.DB
	Cfg *Config
)

type Config struct {
	Mysql *model2.Mysql `json:"mysql"`
	Auth  *model2.Auth  `json:"auth"`
}

type CustomClaims struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	jwt.StandardClaims
}

func CreateToken(name, id string, expiry int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(expiry) * time.Second)
	claims := CustomClaims{
		Name: name,
		ID:   id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(Cfg.Auth.SecretKey))
	return token, err
}

func IsAuthorized(token string) (bool, error) {
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(Cfg.Auth.SecretKey), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractIDFromToken(requestToken string) (string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(Cfg.Auth.SecretKey), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return "", fmt.Errorf("Invalid Token")
	}
	return claims["id"].(string), nil
}

func GetUUID() string {
	return uuid.New().String()
}
