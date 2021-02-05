package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gyan1230/2021/accurics/model/github"
	"github.com/gyan1230/2021/accurics/model/oauth"
)

var (
	//HTTPClient :
	HTTPClient http.Client
	//Tokens : Temporary store in memory database
	Tokens string
)

//OauthRedirectService :
func OauthRedirectService(c *gin.Context) {
	var cid github.IDandSecret
	cid.ClientID = "67762ddbb19d107ae14f"
	cid.Secret = "9c67b48923ea069590dd3f1d2e8d9b0526b4ada1"
	err := c.Request.ParseForm()
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not parse query: %v", err)
		c.JSON(http.StatusBadRequest, "BAD")
	}
	code := c.Request.FormValue("code")
	// Next, lets for the HTTP request to call the github oauth enpoint
	// to get our access token
	reqURL := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", cid.ClientID, cid.Secret, code)
	req, err := http.NewRequest(http.MethodPost, reqURL, nil)
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not create HTTP request: %v", err)
		c.JSON(http.StatusBadRequest, "BAD")
	}

	req.Header.Set("accept", "application/json")

	// Send out the HTTP request
	res, err := HTTPClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not send HTTP request: %v", err)
		c.JSON(http.StatusBadRequest, "BAD")
	}
	defer res.Body.Close()

	// Parse the request body into the `OAuthAccessResponse` struct
	var t oauth.AccessResponse
	if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
		fmt.Fprintf(os.Stdout, "could not parse JSON response: %v", err)
		c.JSON(http.StatusBadRequest, "BAD")
	}
	Tokens = t.AccessToken
	c.Redirect(http.StatusFound, "/welcome.html?access_token="+t.AccessToken)
}
