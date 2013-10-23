package contextio

type Message struct {
	Date           int                    `json:"date"`
	DateIndexed    int                    `json:"date_indexed"`
	DateReceived   int                    `json:"date_received"`
	Addresses      MessageAddressesBlock  `json:"addresses"`
	PersonInfo     MessagePersonInfoBlock `json:"person_info"`
	EmailMessageId string                 `json:"email_message_id"`
	MessageId      string                 `json:"message_id"`
	GmailMessageId string                 `json:"gmail_message_id"`
	GmailThreadId  string                 `json:"gmail_thread_id"`
	Files          []MessageFilesBlock    `json:"files"`
	Subject        string                 `json:"subject"`
	Folders        []string               `json:"folders"`
	Sources        []MessageSourcesBlock  `json:"sources"`
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
	Size               int        `json:"size"`
	Type               string     `json:"type"`
	FileName           string     `json:"file_name"`
	FileNameStructure  [][]string `json:"file_name_structure"`
	BodySection        int        `json:"body_section"`
	FileId             string     `json:"file_id"`
	IsEmbedded         bool       `json:"is_embedded"`
	ContentDisposition string     `json:"content_disposition"`
	ContentId          string     `json:"content_id"`
}

type MessageSourcesBlock struct {
	Label       string `json:"label"`
	ResourceURL string `json:"resource_url"`
}
