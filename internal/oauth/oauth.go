package oauth

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/gitlab"

	"github.com/chuckleheads/builder-api/internal/config"
)

func New(config *config.OauthConfig) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		Endpoint:     findProvider(config.Provider),
	}
}

func findProvider(provider string) oauth2.Endpoint {
	var endpoint oauth2.Endpoint
	// TED TODO: Add more providers, you get the idea
	switch provider {
	case "github":
		endpoint = github.Endpoint
	case "gitlab":
		endpoint = gitlab.Endpoint
	default:
		panic("unknown provider")
	}
	return endpoint
}
