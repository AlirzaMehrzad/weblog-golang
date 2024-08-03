package models

type Comment struct {
	ID     string `json:"id" bson:"_id,omitempty"`
	PostID string `json:"postId" bson:"postId"`
	Body   string `json:"body" bson:"body"`
	UserId string `json:"UserId" bson:"UserId"`
}