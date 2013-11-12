package contextio

type Message struct {
	Date           int                    `json:"date,omitempty"`
	DateIndexed    int                    `json:"date_indexed,omitempty"`
	DateReceived   int                    `json:"date_received,omitempty"`
	Addresses      MessageAddressesBlock  `json:"addresses,omitempty"`
	PersonInfo     MessagePersonInfoBlock `json:"person_info,omitempty"`
	EmailMessageId string                 `json:"email_message_id,omitempty"`
	MessageId      string                 `json:"message_id,omitempty"`
	GmailMessageId string                 `json:"gmail_message_id,omitempty"`
	GmailThreadId  string                 `json:"gmail_thread_id,omitempty"`
	Files          []MessageFilesBlock    `json:"files,omitempty,omitempty"`
	Subject        string                 `json:"subject,omitempty"`
	Folders        []string               `json:"folders,omitempty"`
	Sources        []MessageSourcesBlock  `json:"sources,omitempty"`
	ResourceURL    string                 `json:"resource_url,omitempty"`
}

type MessageAddressesBlock struct {
	From Address   `json:"from"`
	To   []Address `json:"to"`
	CC   []Address `json:"cc"`
}

type Address struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type MessagePersonInfoBlock map[string]ThumbnailBlock

type ThumbnailBlock struct {
	Thumbnail string `json:"thumbnail"`
}

type MessageFilesBlock struct {
	Size               int        `json:"size,omitempty"`
	Type               string     `json:"type"`
	FileName           string     `json:"file_name"`
	FileNameStructure  [][]string `json:"file_name_structure"`
	BodySection        string     `json:"body_section,omitempty"`
	FileId             string     `json:"file_id"`
	IsEmbedded         bool       `json:"is_embedded"`
	ContentDisposition string     `json:"content_disposition"`
	ContentId          string     `json:"content_id"`
}

type MessageSourcesBlock struct {
	Label       string `json:"label"`
	ResourceURL string `json:"resource_url"`
}

type MessageBody struct {
	Type        string `json:"type"`
	Charset     string `json:"charset"`
	Content     string `json:"content"`
	BodySection string `json:"body_section"`
}
