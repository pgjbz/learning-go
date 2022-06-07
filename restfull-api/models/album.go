package models

type Album struct {
	ID     uint64  `json:"id"`     // `json:"id"` is like @JsonProperty(value = "id")
	Title  string  `json:"title"`  // its a go annotation?
	Artist string  `json:"artist"` // get TODO in this to search in future
	Price  float64 `json:"price"`
}
