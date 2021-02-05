package oauth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gyan1230/2021/accurics/model/github"
	"github.com/gyan1230/2021/accurics/services"
)

//GetRedirect :
func GetRedirect(c *gin.Context) {
	services.OauthRedirectService(c)
}

//GetRepo :
func GetRepo(c *gin.Context) {
	var r github.RepoInfo
	err := c.ShouldBindJSON(&r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Not found",
		})
		return
	}
	if r.Owner == "" || r.Repo == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Owner or Repo should not be empty",
		})
		return
	}
	client := services.GetClient(services.Tokens)
	repo, err := services.GetRepo(client, services.Tokens, r.Owner, r.Repo)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"Last commits": repo[0],
	})
}
