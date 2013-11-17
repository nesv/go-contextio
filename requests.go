package contextio

type CreateAccount struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

type CreateSource struct {
	Email                string `json:"email"`
	Server               string `json:"server"`
	Username             string `json:"username"`
	UseSSL               int    `json:"use_ssl"`
	Port                 int    `json:"port"`
	Type                 string `json:"type"`
	OriginIP             string `json:"origin_ip,omitempty"`
	SyncPeriod           string `json:"sync_period,omitempty"`
	ExpungeOnDeletedFlag int    `json:"expunge_on_deleted_flag,omitempty"`
	SyncFlags            int    `json:"sync_flags,omitempty"`
	SyncAllFolders       int    `json:"sync_all_folders,omitempty"`
	RawFileList          int    `json:"raw_file_list,omitempty"`
	Password             string `json:"password,omitempty"`
	ProviderRefreshToken string `json:"password_refresh_token,omitempty"`
	ProviderToken        string `json:"provider_token,omitempty"`
	ProviderTokenSecret  string `json:"provider_token_secret,omitempty"`
	ProviderConsumerKey  string `json:"provider_consumer_key,omitempty"`
	CallbackUrl          string `json:"callback_url,omitempty"`
}

type CreateAccountAndSource struct {
	CreateAccount
	CreateSource
}
