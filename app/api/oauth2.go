package api

import (
	"fmt"
	"github.com/cqqqq777/simple_oauth2/app/dao"
	"github.com/cqqqq777/simple_oauth2/app/model"
	"github.com/cqqqq777/simple_oauth2/app/service"
	"github.com/cqqqq777/simple_oauth2/app/utils"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

func Authorize(c *gin.Context) {
	responseType := c.Query("response_type")
	switch responseType {
	case "code":
		code(c)
	case "token":
		implicit(c)
	default:
		c.JSON(http.StatusOK, gin.H{
			"msg": "invalid format",
		})
	}
}

func code(c *gin.Context) {
	//判断用户是否申请了oauth2服务
	clientID := c.Query("client_id")
	if !utils.IsContain(clientID) {
		c.JSON(http.StatusOK, gin.H{"msg": "Unauthorized Access"})
		return
	}
	//scope表示权限
	//后续开发中完善权限处理
	//scope := c.Param("scope")
	state := c.Query("state")
	redirectUri := c.Query("redirect_uri")
	Code := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000)
	//将客服端id和code都存到redis中在用户申请token时验证
	if err := dao.SetCode(clientID, Code); err != nil {
		c.JSON(http.StatusInternalServerError, "internal error")
		return
	}
	//重定向到用户指定的url
	c.Redirect(http.StatusFound, fmt.Sprintf("%v?code=%06d&state=%v", redirectUri, Code, state))
}

func IssueToken(c *gin.Context) {
	//先验证client_id和client_secret
	clientId := c.GetHeader("client_id")
	secret := c.GetHeader("client_secret")
	if !utils.CheckSecret(clientId, secret) {
		c.JSON(http.StatusOK, "Unauthorized Access")
		return
	}
	param := new(model.Params)
	if err := c.ShouldBindJSON(param); err != nil {
		c.JSON(http.StatusOK, "请求参数错误")
		return
	}
	aToken, rToken, iToken, err := service.IssueToken(param, clientId)
	if err != nil {
		c.JSON(http.StatusOK, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"access_token":  aToken,
		"refresh_token": rToken,
		"scope":         "all",
		"token_type":    "Bearer",
		"expires_in":    1440,
		"id_token":      iToken,
	})
}

func implicit(c *gin.Context) {

}
