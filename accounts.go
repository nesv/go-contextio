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

type Source struct {
	Server             string `json:"server"`
	Label              string `json:"label"`
	Username           string `json:"username"`
	Port               int    `json:"port"`
	AuthenticationType string `json:"authentication_type"`
	Status             string `json:"status"`
	ServiceLevel       string `json:"service_level"`
	SyncPeriod         string `json:"sync_period"`
	UseSSL             bool   `json:"use_ssl"`
	Type               string `json:"type"`
	ResourceURL        string `json:"resource_url"`
}
