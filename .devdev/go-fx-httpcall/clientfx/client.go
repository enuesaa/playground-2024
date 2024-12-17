package clientfx

import (
	"io"
	"net/http"
)

func New() IClient {
	return &Client{}
}

type IClient interface {
	Get(url string) (string, error)
}

type Client struct {}

func (c *Client) Get(url string) (string, error) {
	client := http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodybytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodybytes), nil
}
