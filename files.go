package contextio

type File struct {
	Size               int                    `json:"size"`
	Type               string                 `json:"type"`
	Subject            string                 `json:"subject"`
	Date               int                    `json:"date"`
	DateIndexed        int                    `json:"date_indexed"`
	Addresses          MessageAddressesBlock  `json:"addresses"`
	PersonInfo         MessagePersonInfoBlock `json:"person_info"`
	FileName           string                 `json:"file_name"`
	FileNameStructure  map[string][]string    `json:"file_name_structure"`
	BodySection        int                    `json:"body_section"`
	FileId             string                 `json:"file_id"`
	SupportsPreview    bool                   `json:"supports_preview"`
	IsEmbedded         bool                   `json:"is_embedded"`
	ContentDisposition string                 `json:"content_disposition"`
	ContentId          string                 `json:"content_id"`
	MessageId          string                 `json:"message_id"`
	EmailMessageId     string                 `json:"email_message_id"`
	GmailMessageId     string                 `json:"gmail_message_id"`
	GmailThreadId      string                 `json:"gmail_thread_id"`
}
