package models

//ArticlesByTagAndDate ...
type ArticlesByTagAndDate struct {
	//minItems=1,
	Articles []string `json:"articles"`
	Count    int32    `json:"count,omitempty"`
	//minItems=1,
	RelatedTags []string `json:"related_tags"`
	Tag         string   `json:"tag,omitempty"`
}
