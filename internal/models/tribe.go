package models

type Tribe struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Region     string   `json:"region"`
	Tags       []string `json:"tags"`
	Population int      `json:"population"`
	Counties   int      `json:"counties"`
	Languages  int      `json:"languages"`
}
