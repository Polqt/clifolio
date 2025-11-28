package services

import (
	"context"
	"os"

	"github.com/google/go-github/v79/github"
	"golang.org/x/oauth2"
)

type Repo struct {
	Name        string
	Description string
	HTMLURL     string
	Stars       int
	Language    string
}

func FetchRepos(ctx context.Context, username string) ([]Repo, error) {
	var client *github.Client
	token := os.Getenv("GITHUB_TOKEN")
	if token != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(ctx, ts)
		client = github.NewClient(tc)
	} else {
		client = github.NewClient(nil)
	}

	opts := &github.RepositoryListOptions{
		Type: "all",
		ListOptions: github.ListOptions{
			PerPage: 20,
		},
	}

	var all []*github.Repository
	for {
		repos, res, err := client.Repositories.List(ctx, username, opts)
		if err != nil {
			return nil, err
		}
		all = append(all, repos...)
		if res.NextPage == 0 {
			break
		}
		opts.Page = res.NextPage
	}
	
	out := make([]Repo, 0, len(all))
	for _, r := range all {
		out = append(out, Repo{
			Name: r.GetName(),
			Description: r.GetDescription(),
			HTMLURL: r.GetHTMLURL(),
			Stars: r.GetStargazersCount(),
			Language: r.GetLanguage(),
		})
	}

	return out, nil
}