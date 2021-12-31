package utils

import (
	"strings"
	"time"

	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/config"
	"github.com/WeixinCloud/wxcloudrun-wxcomponent/comm/log"

	"github.com/golang-jwt/jwt/v4"
)

// Claims Claims结构体
type Claims struct {
	UserName string
	jwt.RegisteredClaims
}

// GetToken 获取token
func GetToken(strToken string) string {
	Bearer := "Bearer"
	tokenSlice := strings.Split(strToken, " ")
	if len(tokenSlice) != 2 || tokenSlice[0] != Bearer {
		return ""
	}
	token := tokenSlice[1]
	return token
}

// GenerateToken 生产token
func GenerateToken(id string, username string) (string, error) {
	log.Debugf("jwtExpireTime[%v]", config.ServerConf.JwtExpireTime)
	nowTime := time.Now()
	expiredTime := nowTime.Add(time.Duration(config.ServerConf.JwtExpireTime) * time.Second)
	claims := &Claims{
		UserName: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime), // 过期时间
			IssuedAt:  jwt.NewNumericDate(nowTime),     // 颁发时间
			ID:        id,                              // 编号
			Issuer:    config.ServerConf.JwtIssue,      // 颁发者
			NotBefore: jwt.NewNumericDate(nowTime),     // 生效时间
			Subject:   "User Token",                    // token主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.ServerConf.JwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken 解析token
func ParseToken(token string) (*Claims, error) {
	log.Debugf("token[%s]", token)
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.ServerConf.JwtSecret), nil
	})
	if err != nil {
		log.Error("jwt token fail")
		return nil, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			log.Debug("jwt token succ ", claims)
			return claims, nil
		}
	}
	log.Error("jwt token fail")
	return nil, err
}
