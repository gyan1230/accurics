package github

// Package :
type Package struct {
	FullName      string
	Description   string
	StarsCount    int
	ForksCount    int
	LastUpdatedBy string
}

//IDandSecret :
type IDandSecret struct {
	ClientID string `json:"client_id"`
	Secret   string `json:"secret"`
}

//Authorize :
type Authorize struct {
	ClientID    string `json:"client_id"`
	RedirectURI string `json:"redirect_uri"`
	Secret      string `json:"secret"`
	Login       string `json:"login"`
	Scope       string `json:"scope"`
}

//RepoInfo :
type RepoInfo struct {
	Owner string `json:"owner"`
	Repo  string `json:"repo"`
}
