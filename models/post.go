package models

type Post struct {
	ID    string `json:"id" bson:"_id,omitempty"`
	Title string `json:"title" bson:"title"`
	Text  string `json:"text" bson:"text"`
	Link  string `json:"link" bson:"link"`
	Image string `json:"image" bson:"image"`
}
