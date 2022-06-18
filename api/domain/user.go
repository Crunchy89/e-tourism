package domain

import (
	"context"

	"api/utils/r"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username,omitempty"`
	Email    string             `json:"email" bson:"email,omitempty"`
	Password string             `json:"password" bson:"password,omitempty"`
	IsDelete bool               `json:"is_delete" bson:"is_delete"`
	ISActive bool               `json:"is_active" bson:"is_active"`
	Log      *Log               `json:"log" bson:"log,omitempty"`
}

type UserRepository interface {
	Add(ctx context.Context, d *User) (primitive.ObjectID, r.Ex)
	FetchAll(ctx context.Context) ([]*User, r.Ex)
	FetchByID(ctx context.Context, ID primitive.ObjectID) (*User, r.Ex)
	FetchByUsername(ctx context.Context, username string) (*User, r.Ex)
	FetchByEmail(ctx context.Context, email string) (*User, r.Ex)
	UpdateByID(ctx context.Context, ID primitive.ObjectID, d *User) r.Ex
	UpdateByUsername(ctx context.Context, username string, d *User) r.Ex
	UpdateByEmail(ctx context.Context, email string, d *User) r.Ex
	DeleteByID(ctx context.Context, ID primitive.ObjectID) r.Ex
	ActiveByID(ctx context.Context, ID primitive.ObjectID) r.Ex
}
