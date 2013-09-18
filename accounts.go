package contextio

// Represents a JSON object representing an account, in context.io.
type Account struct {
	Id              string   `json:"id"`
	Username        string   `json:"username"`
	Created         int      `json:"created"`
	Suspended       int      `json:"suspended"`
	EmailAddresses  []string `json:"email_addresses"`
	FirstName       string   `json:"first_name"`
	LastName        string   `json:"last_name"`
	PasswordExpired int      `json:"password_expired"`
	Sources         []Source `json:"sources"`
	NBMessages      int      `json:"nb_messages"`
	NBFiles         int      `json:"nb_files"`
}

type AccountEmailAddressBlock struct {
	Email     string `json:"email"`
	Validated int    `json:"validated"`
	Primary   int    `json:"primary"`
}
