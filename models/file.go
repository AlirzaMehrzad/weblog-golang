package models

type File struct {
	path     string `json:"path" bson:"path"`
	fileName string `json:"fileName" bson:"fileName"`
}