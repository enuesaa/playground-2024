package clifx

import "flag"

func New() ICli {
	cli := Cli{
		Url: "",
	}
	return &cli
}

type ICli interface {
	Launch() error
	GetUrl() string
}

type Cli struct {
	Url string
}

func (c *Cli) Launch() error {
	c.parse()

	if err := c.validate(); err != nil {
		return err
	}
	return nil
}

func (c *Cli) parse() {
	flag.StringVar(&c.Url, "url", "https://example.com/", "Url to call. Example: https://example.com/")
	flag.Parse()
}

func (c *Cli) validate() error {
	return nil
}

func (c *Cli) GetUrl() string {
	return c.Url
}
