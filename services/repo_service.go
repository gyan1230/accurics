package services

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/google/go-github/github"
	git "github.com/gyan1230/2021/accurics/model/github"
	"golang.org/x/oauth2"
)

//GetClient :
func GetClient(token string) *github.Client {
	context := context.Background()
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tokenClient := oauth2.NewClient(context, tokenSource)
	client := github.NewClient(tokenClient)
	return client
}

//GetRepo :
func GetRepo(c *github.Client, token, owner, repo string) ([]*github.RepositoryCommit, error) {
	temprepo, tempresp, err := c.Repositories.Get(context.TODO(), owner, repo)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if tempresp.StatusCode != http.StatusOK {
		log.Println(tempresp.Response)
		return nil, err
	}
	pack := &git.Package{
		FullName:    *temprepo.FullName,
		Description: *temprepo.Description,
		ForksCount:  *temprepo.ForksCount,
		StarsCount:  *temprepo.StargazersCount,
	}
	fmt.Printf("%+v\n", pack)
	commitInfo, _, err := c.Repositories.ListCommits(context.Background(), owner, repo, nil)
	if err != nil {
		fmt.Printf("Problem in commit information %v\n", err)
		return nil, err
	}
	return commitInfo, nil
}
