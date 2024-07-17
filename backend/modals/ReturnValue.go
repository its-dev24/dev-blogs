package modals

import "go.mongodb.org/mongo-driver/bson"

type ReturnValue struct {
	Error error
	Value []bson.M
}
