package callerfx

import (
	"io"
	"net/http"
)

func New() ICaller {
	return &Caller{}
}

type ICaller interface {
	Run(url string) (string, error)
}

type Caller struct {}

func (c *Caller) Run(url string) (string, error) {
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
