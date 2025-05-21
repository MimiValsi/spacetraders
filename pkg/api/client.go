package api

import (
	"context"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	BaseURI *url.URL
	Header *http.Header
	httpClient *http.Client
	//AccountToken string // this is meant to be used only once per week. Not even sure if needed...
}

func NewClient(ctx context.Context, token string) (*Client, error) {
	uri, err := url.Parse("https://api.spacetraders.io/v2/")
	if err != nil {
		return nil, err
	}

	return &Client{
		BaseURI: uri,
		Header: &http.Header{
			"Content-Type": {"application/json"},
			"Authorization": {"Bearer " + token},
		},
		httpClient: &http.Client{
			Timeout: time.Minute,
		},
	}, nil
}


