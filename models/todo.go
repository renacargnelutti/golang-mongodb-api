package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ToDoList struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Task   string             `json:"task"`
	Status bool               `json:"status"`
}
