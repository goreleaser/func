package count

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/google/go-github/v39/github"
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
		"filename:goreleaser language:yaml -path:/vendor",
		&github.SearchOptions{ListOptions: github.ListOptions{}},
	)
	if err != nil {
		return 0, fmt.Errorf("failed to search: %w", err)
	}
	if resp.StatusCode != 200 {
		return 0, errors.New("search request failed")
	}
	return *result.Total, nil
}
