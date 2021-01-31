package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gyan1230/2021/accurics/controller/oauth"
)

func init() {

}

//MapURL :
func mapURL() {
	fs := http.FileServer(http.Dir("public"))
	router.GET("/", gin.WrapH(fs))
	router.GET("/oauth/redirect", oauth.GetRedirect)
	router.POST("/repo", oauth.GetRepo)
}
