package count

import (
	"context"
	"os"

	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

var (
	token = os.Getenv("GITHUB_TOKEN")
	ts    = oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
)

func Count(ctx context.Context) (int, error) {
	var client = github.NewClient(oauth2.NewClient(ctx, ts))
	result, resp, err := client.Search.Code(
		ctx,
		"filename:goreleaser extension:yaml extension:yml path:/",
		&github.SearchOptions{ListOptions: github.ListOptions{}},
	)
	if err != nil {
		return 0, errors.Wrap(err, "failed to search")
	}
	if resp.StatusCode != 200 {
		return 0, errors.New("search request failed")
	}
	return *result.Total, nil
}
