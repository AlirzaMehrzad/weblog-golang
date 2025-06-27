package models

type File struct {
	Path     string `json:"path" bson:"path"`
	FileName string `json:"fileName" bson:"fileName"`
	Section  string `json:"section" bson:"section"`
}
