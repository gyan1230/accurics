package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gyan1230/2021/accurics/controller/oauth"
)

//MapURL :
func mapURL() {
	fs := http.FileServer(http.Dir("public"))
	ginRouter.StaticFile("/welcome.html", "./public/welcome.html")
	ginRouter.GET("/", gin.WrapH(fs))
	ginRouter.GET("/oauth/redirect", oauth.GetRedirect)
	ginRouter.POST("/repo", oauth.GetRepo)
}
