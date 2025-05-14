package utils

import (
	"fmt"
	"io"
)

type UpdateResult int

const (
	NoChange UpdateResult = iota
	Updated
	Failed
)

type DynDnsConfig struct {
	DynDnsDomain         string
	DynDnsUpdateEndpoint string
	DynDnsToken          string
	CurrentIp            string
	RegisteredIp         string
	Client               HttpClient
}

func (c *DynDnsConfig) check() bool {
	return c.CurrentIp != c.RegisteredIp
}

func (c *DynDnsConfig) update() (string, error) {
	dynDnsUpdateEndpointUrl := fmt.Sprintf("%s?domains=%s&token=%s[&ip=%s]", c.DynDnsUpdateEndpoint, c.DynDnsDomain, c.DynDnsToken, c.CurrentIp)

	resp, err := c.Client.Get(dynDnsUpdateEndpointUrl)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func CheckAndUpdate(c *DynDnsConfig) (UpdateResult, error) {
	if c.check() {
		_, err := c.update()
		if err != nil {
			return Failed, err
		}
		return Updated, nil
	}
	return NoChange, nil
}
