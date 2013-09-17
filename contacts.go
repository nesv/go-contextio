package contextio

type ContactList struct {
	Query   ContactListQueryBlock `json:"query"`
	Matches []Contact             `json:"matches"`
}

type ContactListQueryBlock struct {
	Limit        int    `json:"limit"`
	Offset       int    `json:"offset"`
	ActiveAfter  int    `json:"active_after"`
	ActiveBefore int    `json:"active_before"`
	Search       string `json:"search"`
}

type Contact struct {
	Email     string `json:"email"`
	Count     int    `json:"count"`
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
}
