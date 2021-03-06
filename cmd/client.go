package cmd

import (
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/tsuru/config"
	"io/ioutil"
	"net/http"
	"net/url"
	"github.com/megamsys/libgo/cmd"
)

type Client struct {
	HTTPClient     *http.Client
	context        *Context
	Authly         *cmd.Authly
	progname       string
	currentVersion string
	versionHeader  string
}

func NewClient(client *http.Client, context *Context, manager *Manager) *Client {
	return &Client{
		HTTPClient:     client,
		context:        context,
		Authly:         &cmd.Authly{},
		progname:       manager.name,
		currentVersion: manager.version,
		versionHeader:  manager.versionHeader,
	}
}

func (c *Client) detectClientError(err error) error {
	urlErr, ok := err.(*url.Error)
	if !ok {
		return err
	}
	switch urlErr.Err.(type) {
	case x509.UnknownAuthorityError:
		target, _ := config.GetString("api:host")
		return fmt.Errorf("Failed to connect to api server (%s): %s", target, urlErr.Err)
	}
	target, _ := config.GetString("api:host")
	return fmt.Errorf("Failed to connect to api server (%s): %s.", target, urlErr.Err)
}

func (c *Client) Do(request *http.Request) (*http.Response, error) {
	for headerKey, headerVal := range c.Authly.AuthMap {
		request.Header.Add(headerKey, headerVal)
	}

	request.Close = true
	response, err := c.HTTPClient.Do(request)
	err = c.detectClientError(err)
	if err != nil {
		return nil, err
	}
	supported := response.Header.Get(c.versionHeader)
	format := `################################################################

WARNING: You're using an unsupported version of %s.

You must have at least version %s, your current
version is %s.

Please go to http://docs.tsuru.io/en/latest/install/client.html
and download the last version.

################################################################

`
	if !validateVersion(supported, c.currentVersion) {
		fmt.Fprintf(c.context.Stderr, format, c.progname, supported, c.currentVersion)
	}
	if response.StatusCode > 399 {
		defer response.Body.Close()
		result, _ := ioutil.ReadAll(response.Body)
		return response, errors.New(string(result))
	}
	return response, nil

	
}
