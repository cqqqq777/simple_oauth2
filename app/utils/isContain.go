package utils

import g "github.com/cqqqq777/simple_oauth2/app/global"

func IsContain(id string) bool {
	for _, v := range g.Config.Oauth2.Client {
		if id == v.ClientID {
			return true
		}
	}
	return false
}

func CheckSecret(id, secret string) bool {
	for _, v := range g.Config.Oauth2.Client {
		if id == v.ClientID {
			if secret == v.ClientSecret {
				return true
			}
		}
	}
	return false
}
