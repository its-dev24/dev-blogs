package modals

import "go.mongodb.org/mongo-driver/bson/primitive"

type Blog struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty"`
	Title    string             `json:"title,omitempty" bson:"title,omitmepty"`
	Author   string             `json:"Author,omitempty" bson:"Author,omitempty"`
	BlogBody string             `json:"Blog,omitempty" bson"Author,omitempty"`
}
