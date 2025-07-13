package trello

import (
	"net/http"
	"time"
)

type TrelloClient struct {
	HTTPClient http.Client
	APIKey     string
	APIToken   string
	Endpoint   string
}

func NewTrelloClient(apiKey, apiToken, endpoint string) *TrelloClient {
	return &TrelloClient{
		HTTPClient: http.Client{
			Timeout: 10 * time.Second,
		},
		APIToken: apiToken,
		APIKey:   apiKey,
		Endpoint: endpoint,
	}
}
