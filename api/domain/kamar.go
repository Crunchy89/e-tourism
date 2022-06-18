package domain

import (
	"api/utils/r"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Kamar struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	KamarID     primitive.ObjectID `json:"kamar_id" bson:"kamar_id,omitempty"`
	JumlahKamar uint16             `json:"jumlah_kamar" bson:"jumlah_kamar,omitempty"`
	Harga       uint64             `json:"harga" bson:"harga,omitempty"`
	Diskon      uint64             `json:"diskon" bson:"diskon,omitempty"`
	IsDelete    bool               `json:"is_delete" bson:"is_delete"`
	IsActive    bool               `json:"is_active" bson:"is_active"`
	Log         *Log               `json:"log" bson:"log,omitempty"`
}

type KamarRepository interface {
	Add(ctx context.Context, d *Kamar) (primitive.ObjectID, r.Ex)
	FetchAll(ctx context.Context) ([]*Kamar, r.Ex)
	FetchByID(ctx context.Context, ID primitive.ObjectID) (*Kamar, r.Ex)
	UpdateByID(ctx context.Context, ID primitive.ObjectID, d *Kamar) r.Ex
	DeleteByID(ctx context.Context, ID primitive.ObjectID) r.Ex
	Active(ctx context.Context, ID primitive.ObjectID) r.Ex
}
