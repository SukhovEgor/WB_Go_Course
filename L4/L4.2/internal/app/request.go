package app

type GrepRequest struct {
	Pattern    string `json:"pattern"`
	File       string `json:"file"`
	IgnoreCase bool   `json:"ignore_case"`
	Invert     bool   `json:"invert"`
	CountOnly  bool   `json:"count_only"`
}