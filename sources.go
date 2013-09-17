package contextio

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
