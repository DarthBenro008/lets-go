package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID     primitive.ObjectID `json:"id,omitempty"  bson:"_id,omitempty"`
	Title  string             `json:"title" binding:"required,min=2" bson:"title"`
	Author string             `json:"author" bson:"author,omitempty"`
}
