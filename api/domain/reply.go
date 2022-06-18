package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Reply struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ReviewID primitive.ObjectID `json:"review_id" bson:"review_id,omitempty"`
	UserID   primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	Comment  string             `json:"comment" bson:"comment,omitempty"`
	IsDelete bool               `json:"is_delete" bson:"is_delete"`
	Log      *Log               `json:"log" bson:"log,omitempty"`
}
