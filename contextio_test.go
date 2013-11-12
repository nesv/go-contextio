package contextio

import (
	"log"
	"os"
	"testing"
)

var (
	testAccount  Account
	testKey      string
	testSecret   string
	testMessages []Message
)

func init() {
	testKey = os.Getenv("CONTEXTIOKEY")
	if len(testKey) == 0 {
		log.Fatal("CONTEXTIOKEY is not defined")
	}

	testSecret = os.Getenv("CONTEXTIOSECRET")
	if len(testSecret) == 0 {
		log.Fatal("CONTEXTIOSECRET is not defined")
	}
}

func newTestClient() *Client {
	c := New(testKey, testSecret)
	c.EnableDebugging(true)
	return c
}

func TestClient(t *testing.T) {
	client := newTestClient()
	resp, err := client.Do("GET", nil, "accounts")
	if err != nil {
		t.Fatal(err)
	}

	var accounts []Account
	if err = client.Unmarshal(resp, &accounts); err != nil {
		t.Fatal(err)
	}
	for i, account := range accounts {
		t.Log("Account:", i, account.Username)
	}

	testAccount = accounts[0]
}

func TestMessages(t *testing.T) {
	client := newTestClient()
	resp, err := client.Do("GET", nil, "accounts", testAccount.Id, "messages")
	if err != nil {
		t.Fatal(err)
	}

	var messages []Message
	if err = client.Unmarshal(resp, &messages); err != nil {
		t.Fatal(err)
	}

	for _, m := range messages {
		t.Log("Retrieved message", m.MessageId)
	}

	t.Log(len(messages), "messages total")
	testMessages = messages
}

func TestRetrieveMessageBody(t *testing.T) {
	client := newTestClient()
	for _, m := range testMessages {
		t.Log("Fetching message:", m.MessageId)
		params := map[string]string{"type": "text/plain"}
		resp, err := client.Do("GET", params, "accounts", testAccount.Id, "messages", m.MessageId, "body")
		//resp, err := client.Do("GET", nil, "accounts", testAccount.Id, "messages", m.MessageId, "body")
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%#v", resp)
		/*if resp.ContentLength <= 0 {
			t.Error("Skipping message due to no content in the response body...")
			continue
		}*/
		//var body MessageBody
		var body []MessageBody
		if err = client.Unmarshal(resp, &body); err != nil {
			t.Error(err.Error())
		}
		//t.Log("Message body:", body.Content)
		/*for _, b := range body {
			t.Logf("Message body (%s):\n%s", b.Type, b.Content)
		}*/
		t.Logf("Message body (%s):\n%s", body[0].Type, body[0].Content)
	}
}
