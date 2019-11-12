package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	IsCompleted bool               `json:"isCompleted"`
}
