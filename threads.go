package contextio

type ThreadList []string

type Thread struct {
	GmailThreadId   string                 `json:"gmail_thread_id"`
	EmailMessagesId []string               `json:"email_messages_id"`
	PersonInfo      MessagePersonInfoBlock `json:"person_info"`
	Messages        []Message              `json:"messages"`
}
