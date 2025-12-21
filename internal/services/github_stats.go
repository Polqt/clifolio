package services

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/go-github/v79/github"
	"golang.org/x/oauth2"
)

type GitHubStats struct {
	TotalRepos   int
	TotalStars   int
	TotalForks   int
	PublicGists  int
	Followers    int
	Following    int
	TotalCommits int
	UpdatedAt    time.Time
}

func FetchGitHubStats(ctx context.Context, username string) (*GitHubStats, error) {
	var client *github.Client
	token := os.Getenv("GITHUB_TOKEN")

	if token != "" {
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
		tc := oauth2.NewClient(ctx, ts)
		client = github.NewClient(tc)
	} else {
		client = github.NewClient(nil)
	}

	user, _, err := client.Users.Get(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch user: %w", err)
	}

	repos, err := FetchRepos(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch repos: %w", err)
	}

	totalStars := 0
	totalForks := 0

	for _, repo := range repos {
		totalStars += repo.Stars
	}

	stats := &GitHubStats{
		TotalRepos: user.GetPublicRepos(),
		TotalStars: totalStars,
		TotalForks: totalForks,
		PublicGists: user.GetPublicGists(),
		Followers: user.GetFollowers(),
		Following: user.GetFollowing(),
		UpdatedAt: time.Now(),
	}

	return stats, nil
}