package utils

import (
	"errors"
	g "github.com/cqqqq777/simple_oauth2/app/global"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type (
	MyClaim struct {
		Id int64
		jwt.RegisteredClaims
	}
	IdClaim struct {
		Nickname string
		jwt.RegisteredClaims
	}
)

func GetToken(id int64) (aToken, rToken string, err error) {
	claim := MyClaim{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Truncate(time.Second)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(g.Config.Jwt.ExpiresTime) * time.Second)),
		},
	}
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(g.Config.Jwt.SecretKey))
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 720)),
	},
	).SignedString([]byte(g.Config.Jwt.SecretKey))
	return
}

func ParseToken(tokenStr string) (claim *MyClaim, err error) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(g.Config.Jwt.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*MyClaim)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// RefreshToken 刷新token
func RefreshToken(aToken, rToken string) (NewAToken, RToken string, err error) {
	//判断refresh token是否有效，无效直接返回
	if _, err = jwt.Parse(rToken, KeyFunc()); err != nil {
		return
	}
	//判断access token是否是过期错误并从中解析数据
	var claim MyClaim
	_, err = jwt.ParseWithClaims(aToken, &claim, KeyFunc())
	v, _ := err.(jwt.ValidationError)
	if v.Errors == jwt.ValidationErrorExpired {
		// 生成新的access token，但refresh token不变
		NewAToken, _, err = GetToken(claim.Id)
		return NewAToken, rToken, err
	}
	return
}

func KeyFunc() func(token *jwt.Token) (interface{}, error) {
	return func(t *jwt.Token) (interface{}, error) {
		return []byte(g.Config.Jwt.SecretKey), nil
	}
}

func GetIDToken(nickname string) (idToken string, err error) {
	claim := IdClaim{
		Nickname: nickname,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Truncate(time.Second)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(g.Config.Jwt.ExpiresTime) * time.Second)),
		},
	}
	idToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(g.Config.Jwt.SecretKey))
	return
}
