package entities

type Post struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}
