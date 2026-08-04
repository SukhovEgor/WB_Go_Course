package app

type GrepResponse struct {
	Lines []string `json:"lines"`
	Count int      `json:"count"`
	Node  string   `json:"node"`
	Error string   `json:"error,omitempty"`
}