package domain

import (
	"api/utils/r"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Penginapan struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Bintang     uint8              `json:"bintang" bson:"bintang,omitempty"`
	NamaHotel   string             `json:"nama_hotel" bson:"nama_hotel,omitempty"`
	Slug        string             `json:"slug" bson:"slug,omitempty"`
	Alamat      string             `json:"alamat" bson:"alamat,omitempty"`
	NoHp        string             `json:"no_hp" bson:"no_hp,omitempty"`
	Jenis       string             `json:"jenis" bson:"jenis,omitempty"`
	TotalKamar  uint16             `json:"total_kamar" bson:"total_kamar,omitempty"`
	MediaSosial *MediaSosial       `json:"media_sosial" bson:"media_sosial,omitempty"`
	Koordinat   *Koordinat         `json:"koordinat" bson:"koordinat,omitempty"`
	Rating      float32            `json:"rating" bson:"rating,omitempty"`
	TotalRating uint16             `json:"total_rating" bson:"total_rating,omitempty"`
	TotalReview uint32             `json:"review" bson:"review,omitempty"`
	IsDelete    bool               `json:"is_delete" bson:"is_delete"`
	IsActive    bool               `json:"is_active" bson:"is_active"`
	Log         *Log               `json:"log" bson:"log,omitempty"`
}
type PenginapanPagination struct {
	Data  []*Penginapan `json:"data"`
	Page  int64         `json:"page"`
	Limit int64         `json:"limit"`
	Total int64         `json:"total"`
}

type PenginapanRepository interface {
	Add(ctx context.Context, d *Penginapan) (primitive.ObjectID, r.Ex)
	FetchAll(ctx context.Context) ([]*Penginapan, r.Ex)
	FetchByID(ctx context.Context, ID primitive.ObjectID) (*Penginapan, r.Ex)
	UpdateByID(ctx context.Context, ID primitive.ObjectID, d *Penginapan) r.Ex
	DeleteByID(ctx context.Context, ID primitive.ObjectID) r.Ex
	Active(ctx context.Context, ID primitive.ObjectID) r.Ex
	FetchBySlug(ctx context.Context, slug string) (*Penginapan, r.Ex)
	Pagination(ctx context.Context, page int64, limit int64) ([]*Penginapan, r.Ex)
	Search(ctx context.Context, page int64, limit int64, search string) ([]*Penginapan, r.Ex)
	CountAll(ctx context.Context) (int64, r.Ex)
}
