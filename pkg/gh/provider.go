package gh

import (
	"context"
	"github.com/google/go-github/v21/github"
	"golang.org/x/oauth2"
)

type Provider struct {
	client *github.Client
}

func New(accessToken string) *Provider {
	ts := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: accessToken,
	})
	return &Provider{
		client: github.NewClient(oauth2.NewClient(context.Background(), ts)),
	}
}


func (p *Provider) Checks() *github.ChecksService {
	return p.client.Checks
}
