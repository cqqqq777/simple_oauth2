package boot

import (
	"github.com/cqqqq777/simple_oauth2/app/api"
	"github.com/gin-gonic/gin"
)

func RoutersInit() {
	r := gin.Default()
	r.GET("/authorize", api.Authorize)
	r.POST("/token", api.IssueToken)
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
