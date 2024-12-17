package cmdfx

import (
	"flag"
	"fmt"

	"github.com/enuesaa/.devdev/go-fx-httpcall/clientfx"
)

func New(client clientfx.IClient) ICmd {
	cmd := Cmd{
		client: client,
		url: "",
	}
	return &cmd
}

type ICmd interface {
	Parse() error
	Run() error
}

type Cmd struct {
	client clientfx.IClient

	url string
}

func (c *Cmd) Parse() error {
	flag.StringVar(&c.url, "url", "", "Url to call. Example: https://example.com/")
	flag.Parse()

	// validate
	if c.url == "" {
		return fmt.Errorf("missing required flag: -url")
	}

	return nil
}

func (c *Cmd) Run() error {
	// http リクエスト
	body, err := c.client.Get(c.url)
	if err != nil {
		return err
	}
	fmt.Printf("%s", body)

	return nil
}
