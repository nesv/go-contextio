package contextio

type ConnectToken struct {
	Token       string                   `json:"token"`
	Email       string                   `json:"email"`
	Created     int                      `json:"created"`
	Used        int                      `json:"used"`
	Expires     interface{}              `json:"expires"` // Can be an int, or a boolean
	CallbackUrl string                   `json:"callback_url"`
	FirstName   string                   `json:"first_name"`
	LastName    string                   `json:"last_name"`
	Account     ConnectTokenAccountBlock `json:"account,omitempty"`
}

type ConnectTokenAccountBlock struct {
	Id             string                     `json:"id"`
	Created        int                        `json:"created"`
	Suspended      int                        `json:"suspended"`
	EmailAddresses []AccountEmailAddressBlock `json:"email_addresses"`
	FirstName      string                     `json:"first_name"`
	LastName       string                     `json:"last_name"`
	Sources        []Source                   `json:"sources"`
}
