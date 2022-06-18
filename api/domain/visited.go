package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Visited struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ForeignID primitive.ObjectID `json:"foreign_id" bson:"foreign_id,omitempty"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	IsDelete  bool               `json:"is_delete" bson:"is_delete"`
	Log       *Log               `json:"log" bson:"log,omitempty"`
}
