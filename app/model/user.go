package model

import (
	"errors"
	g "github.com/cqqqq777/simple_oauth2/app/global"
	"github.com/cqqqq777/simple_oauth2/app/utils"
)

const CtxGetUID = "UserID"

type (
	User interface {
		Login() (string, error)
		GetNickName() (string, error)
	}
	UserInfo struct {
		Id       int64  `json:"id" db:"id"`
		Username string `json:"username" db:"username"`
		Password string `json:"password" db:"password"`
		NickName string `json:"nickname" db:"nickname"`
	}
)

var (
	ErrorWrongPassword = errors.New("wrong password")
)

func (u *UserInfo) Login() (aToken, rToken string, err error) {
	var password string
	if err = g.Mdb.QueryRow("select password from users where username = ? ", u.Username).Scan(&password); err != nil {
		return "", "", err
	}
	if password != u.Password {
		return "", "", ErrorWrongPassword
	}
	if err = g.Mdb.QueryRow("select id from users where username = ?", u.Username).Scan(&u.Id); err != nil {
		return "", "", err
	}
	aToken, rToken, err = utils.GetToken(u.Id)
	if err != nil {
		return
	}
	return
}

func (u *UserInfo) GetNickName() error {
	err := g.Mdb.QueryRow("select nickname from users where id = ?", u.Id).Scan(&u.NickName)
	return err
}
