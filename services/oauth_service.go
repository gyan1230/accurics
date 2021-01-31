package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gyan1230/2021/accurics/model/oauth"
)

//RedirectSvc :
func RedirectSvc(c *gin.Context) oauth.AccessResponse {
	var cid oauth.IDandSecret
	cid.ClientID = "67762ddbb19d107ae14f"
	cid.Secret = "9c67b48923ea069590dd3f1d2e8d9b0526b4ada1"
	httpClient := http.Client{}
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
	res, err := httpClient.Do(req)
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

	// Finally, send a response to redirect the user to the "welcome" page
	// with the access token
	c.Writer.Header().Set("Location", "/welcome.html?access_token="+t.AccessToken)
	return t
}
