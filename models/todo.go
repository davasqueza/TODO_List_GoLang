package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	title       string
	description string
	isCompleted bool
}
