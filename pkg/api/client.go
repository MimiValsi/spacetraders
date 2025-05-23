package api

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/MimiValsi/spacetraders/internal/database"
)

type Client struct {
	BaseURI    *url.URL
	Header     *http.Header
	HttpClient *http.Client

	DB *database.Queries
}

func NewClient(ctx context.Context, token string, db *database.Queries) (*Client, error) {
	uri, err := url.Parse("https://api.spacetraders.io/v2/")
	if err != nil {
		return nil, err
	}

	return &Client{
		BaseURI: uri,
		Header: &http.Header{
			"Content-Type":  {"application/json"},
			"Authorization": {"Bearer " + token},
		},
		HttpClient: &http.Client{
			Timeout: time.Minute,
		},
		DB: db,
	}, nil
}
