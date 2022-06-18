package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Fasilitas struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ForeignKey    primitive.ObjectID `json:"foreign_key" bson:"foreign_key,omitempty"`
	NamaFasilitas string             `json:"nama_fasilitas" bson:"nama_fasilitas,omitempty"`
	Type          string             `json:"type" bson:"type,omitempty"`
	IsDelete      bool               `json:"is_delete" bson:"is_delete"`
	IsActive      bool               `json:"is_active" bson:"is_active"`
	Log           *Log               `json:"log" bson:"log,omitempty"`
}
