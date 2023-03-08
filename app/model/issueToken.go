package model

type Params struct {
	Code        int64  `json:"code" binding:"required"`
	GrandType   string `json:"grand_type"`
	RedirectUri string `json:"redirect_uri" binding:"required"`
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
}
