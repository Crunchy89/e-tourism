package domain

import (
	"api/utils/r"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Fasilitas struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ForeignID     primitive.ObjectID `json:"foreign_id" bson:"foreign_id,omitempty"`
	NamaFasilitas string             `json:"nama_fasilitas" bson:"nama_fasilitas,omitempty"`
	Jumlah        uint64             `json:"jumlah" bson:"jumlah,omitempty"`
	IsDelete      bool               `json:"is_delete" bson:"is_delete"`
	IsActive      bool               `json:"is_active" bson:"is_active"`
	Log           *Log               `json:"log" bson:"log,omitempty"`
}

type FasilitasRepository interface {
	Add(ctx context.Context, d *Fasilitas) (primitive.ObjectID, r.Ex)
	FetchAll(ctx context.Context) ([]*Fasilitas, r.Ex)
	FetchByID(ctx context.Context, ID primitive.ObjectID) (*Fasilitas, r.Ex)
	UpdateByID(ctx context.Context, ID primitive.ObjectID, d *Fasilitas) r.Ex
	DeleteByID(ctx context.Context, ID primitive.ObjectID) r.Ex
	Active(ctx context.Context, ID primitive.ObjectID) r.Ex
	FetchByForeignID(ctx context.Context, foreignID primitive.ObjectID) ([]*Fasilitas, r.Ex)
	PaginationByForeignID(ctx context.Context, foreignID primitive.ObjectID, page int64, limit int64) ([]*Fasilitas, r.Ex)
	SearchByForeignID(ctx context.Context, foreignID primitive.ObjectID, page int64, limit int64, keyword string) ([]*Fasilitas, r.Ex)
}
