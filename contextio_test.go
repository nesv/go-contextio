package contextio

import (
	"os"
	"testing"
)

func TestClient(t *testing.T) {
	key := os.Getenv("CONTEXTIOKEY")
	if len(key) == 0 {
		t.Fatal("CONTEXTIOKEY is not defined")
	}

	secret := os.Getenv("CONTEXTIOSECRET")
	if len(secret) == 0 {
		t.Fatal("CONTEXTIOSECRET is not defined")
	}

	client := New(key, secret)
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
}
