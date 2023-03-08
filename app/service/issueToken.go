package service

import (
	"github.com/cqqqq777/simple_oauth2/app/dao"
	"github.com/cqqqq777/simple_oauth2/app/model"
	"github.com/cqqqq777/simple_oauth2/app/utils"
	"log"
)

func IssueToken(params *model.Params, id string) (aToken, rToken, iToken string, err error) {
	code, err := dao.GetCode(id)
	if err != nil {
		return "", "", "", err
	}
	if params.Code != code {
		return "", "", "", err
	}
	user := &model.UserInfo{
		Username: params.Username,
		Password: params.Password,
	}
	//获取access token和refresh token
	aToken, rToken, err = user.Login()
	if err != nil {
		log.Println(err)
		return "", "", "", err
	}
	//获取id token
	if err = user.GetNickName(); err != nil {
		return "", "", "", err
	}
	iToken, err = utils.GetIDToken(user.NickName)
	if err != nil {
		return "", "", "", err
	}
	return
}
