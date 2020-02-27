package models

//Article ...
type Article struct {
	Body string `json:"body,omitempty"`
	Date string `json:"date"`
	ID   string `json:"id,omitempty"`
	//minItems=1,
	Tags  []string `json:"tags"`
	Title string   `json:"title,omitempty"`
}
