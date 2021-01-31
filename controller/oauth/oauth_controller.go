package oauth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gyan1230/2021/accurics/model/oauth"
	"github.com/gyan1230/2021/accurics/services"
)

//Tokens : Temporary store in memory database
var Tokens = make(map[string]string)

//GetRedirect :
func GetRedirect(c *gin.Context) {
	Tokens = make(map[string]string)
	token := services.RedirectSvc(c)
	Tokens["User1"] = token.AccessToken
}

//GetRepo :
func GetRepo(c *gin.Context) {
	var r oauth.RepoInfo
	err := c.ShouldBindJSON(&r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Not found",
		})
	}
	token := Tokens["User1"]
	client := services.GetClient(token)
	repo, err := services.GetRepo(client, token, r.Owner, r.Repo)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusAccepted, gin.H{
		"Last commits": repo[0],
	})
}
