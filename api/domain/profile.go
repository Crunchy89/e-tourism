package domain

import (
	"api/utils/r"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Profile struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID   primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	FullName string             `json:"full_name" bson:"full_name,omitempty"`
	Alamat   string             `json:"alamat" bson:"alamat,omitempty"`
	Profile  string             `json:"avatar" bson:"avatar,omitempty"`
	IsDelete bool               `json:"is_delete" bson:"is_delete"`
	Log      *Log               `json:"log" bson:"log,omitempty"`
}

type ProfileRepository interface {
	Add(ctx context.Context, d *Profile) (primitive.ObjectID, r.Ex)
	FetchAll(ctx context.Context) ([]*Profile, r.Ex)
	FetchByID(ctx context.Context, ID primitive.ObjectID) (*Profile, r.Ex)
	UpdateByID(ctx context.Context, ID primitive.ObjectID, d *Profile) r.Ex
	DeleteByID(ctx context.Context, ID primitive.ObjectID) r.Ex
}
