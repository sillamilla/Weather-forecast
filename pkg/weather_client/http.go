package weather_client

import (
	"io"
	"net/http"
	"time"
)

type Client struct {
	http     *http.Client
	apiUrl   string
	apiToken string
}

func NewClient(apiUrl string, apiToken string) *Client {
	return &Client{
		http: &http.Client{
			Timeout: time.Second * 10,
		},
		apiUrl:   apiUrl,
		apiToken: apiToken,
	}
}

func (c *Client) get(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	do, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer do.Body.Close()

	body, err := io.ReadAll(do.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
