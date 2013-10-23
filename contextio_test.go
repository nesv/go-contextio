package contextio

import (
	"log"
	"os"
	"testing"
)

var (
	testAccount Account
	testClient  *Client
	testKey     string
	testSecret  string
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

func TestClient(t *testing.T) {
	testClient = New(testKey, testSecret)
	resp, err := testClient.Do("GET", nil, "accounts")
	if err != nil {
		t.Fatal(err)
	}

	var accounts []Account
	if err = testClient.Unmarshal(resp, &accounts); err != nil {
		t.Fatal(err)
	}
	for i, account := range accounts {
		t.Log("Account:", i, account.Username)
	}

	testAccount = accounts[0]
}

func TestMessages(t *testing.T) {
	resp, err := testClient.Do("GET", nil, "accounts", testAccount.Id, "messages")
	if err != nil {
		t.Fatal(err)
	}

	var messages []Message
	if err = testClient.Unmarshal(resp, &messages); err != nil {
		t.Fatal(err)
	}

	t.Log("Retrieved", len(messages), "messages")
}
