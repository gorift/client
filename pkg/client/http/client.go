package http

import (
	"fmt"
	"net/http"

	"github.com/gorift/gorift/pkg/balance"
)

type Client struct {
	rawClient *http.Client
	balancer  *balance.Balancer
}

func New() *Client {
	return &Client{}
}

func (c *Client) Get(endpoint string) (*http.Response, error) {
	picked, err := c.balancer.Pick()
	if err != nil {
		return nil, err
	}
	url := generateURL(picked.Address.String(), endpoint)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func generateURL(address, endpoint string) string {
	return fmt.Sprintf("%s%s", address, endpoint)
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.rawClient.Do(req)
}
