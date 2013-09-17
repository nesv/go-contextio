package contextio

type Error struct {
	Type  string `json:"type"`
	Code  int    `json:"code"`
	Value string `json:"value"`
}
