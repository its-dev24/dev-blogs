package modals

import "go.mongodb.org/mongo-driver/bson/primitive"

type Blog struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title    string             `json:"title,omitempty" bson:"title,omitmepty"`
	Author   string             `json:"author,omitempty" bson:"author,omitempty"`
	BlogBody string             `json:"blog,omitempty" bson:"blog,omitempty"`
}
