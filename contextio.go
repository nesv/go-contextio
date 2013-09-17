//
// A client-side library for interacting with Context.IO.
//
package contextio

import (
	"encoding/json"
	"fmt"
	"github.com/mrjones/oauth"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	API_URL     = "https://api.context.io"
	API_VERSION = "2.0"
)

func New(key, secret string) *Client {
	oc := oauth.NewConsumer(key, secret, nil)
	return &Client{consumer: oc}
}

type Client struct {
	consumer *oauth.Consumer
}

func (c *Client) Do(method string, parts ...string) (resp *http.Response, err error) {
	// Build the URL.
	url := fmt.Sprintf("%s/%s/%s", API_URL, API_VERSION, strings.Join(parts, "/"))

	// Call the appropriate function on the consumer, depending on the provided
	// value of `method`.
	switch strings.ToUpper(method) {
	case "GET":
		resp, err = c.consumer.Get(url, nil, nil)

	case "POST":
		resp, err = c.consumer.Post(url, nil, nil)

	case "PUT":
		resp, err = c.consumer.Put(url, nil, nil)

	case "DELETE":
		resp, err = c.consumer.Delete(url, nil, nil)
	}

	return
}

func (c *Client) Unmarshal(f func() (*http.Response, error), v interface{}) error {
	resp, err := f()
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, v)
}
