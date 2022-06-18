package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Travel struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	NamaTravel  string             `json:"nama_travel" bson:"nama_travel,omitempty"`
	Pemilik     string             `json:"pemilik" bson:"pemilik,omitempty"`
	Alamat      string             `json:"alamat" bson:"alamat,omitempty"`
	NoHp        string             `json:"no_hp" bson:"no_hp,omitempty"`
	Email       string             `json:"email" bson:"email,omitempty"`
	Koordinat   *Koordinat         `json:"koordinat" bson:"koordinat,omitempty"`
	MediaSosial *MediaSosial       `json:"media_sosial" bson:"media_sosial,omitempty"`
	IsDelete    bool               `json:"is_delete" bson:"is_delete"`
	IsActive    bool               `json:"is_active" bson:"is_active"`
	Log         *Log               `json:"log" bson:"log,omitempty"`
}
