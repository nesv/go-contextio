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
	oc := oauth.NewConsumer(key, secret, oauth.ServiceProvider{})
	return &Client{consumer: oc, key: key, secret: secret}
}

type Client struct {
	consumer *oauth.Consumer
	key      string
	secret   string
}

/*
Enables debugging of the OAuth consumer when set to true. Providing a value of
false disables debugging.
*/
func (c *Client) EnableDebugging(debug bool) {
	c.consumer.Debug(debug)
}

/*
Perform a two-legged OAuth 1.0 request against the context.io API.

The supported values for the method are "GET", "POST", "PUT", and "DELETE".

The params string-string map is for any extra parameters you wish to provide to
an endpoint, that will be URL-encoded.

The rest of the arguments are the various parts of the API endpoint's URI, done
like a print-like function:

	var account Account
	...
	c.Do("GET", nil, "accounts", account.Id, "files")

...this is so that you do not have to constantly re-create the a URI string to
pass to this function.
*/
func (c *Client) Do(method string, params map[string]string, parts ...string) (resp *http.Response, err error) {
	// Build the URL.
	url := fmt.Sprintf("%s/%s/%s", API_URL, API_VERSION, strings.Join(parts, "/"))

	// The access token
	accessToken := &oauth.AccessToken{}

	// Call the appropriate function on the consumer, depending on the provided
	// value of `method`.
	switch strings.ToUpper(method) {
	case "GET":
		resp, err = c.consumer.Get(url, params, accessToken)

	case "POST":
		resp, err = c.consumer.Post(url, params, accessToken)

	case "PUT":
		resp, err = c.consumer.Put(url, "", params, accessToken)

	case "DELETE":
		resp, err = c.consumer.Delete(url, params, accessToken)
	}

	return
}

/*
Unmarshals the body of the provided HTTP response, into the provided type.

For example, if you called:

	response, err := c.Do("GET", "accounts")

...you would call *c.Unmarshal() like so:

	var accounts []contextio.Account
	err = c.Unmarshal(response, &accounts)
*/
func (c *Client) Unmarshal(response *http.Response, v interface{}) error {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}

/*
 Copies the details of the existing client into a the new one.
*/
func (c *Client) CopyTo(newclient *Client) {
	*newclient = *c
}
