package domain

import (
	"api/utils/r"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Wisata struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	NamaObjekWisata string             `json:"nama_objek_wisata" bson:"nama_objek_wisata,omitempty"`
	Slug            string             `json:"slug" bson:"slug,omitempty"`
	PotensiWisata   string             `json:"potensi_wisata" bson:"potensi_wisata,omitempty"`
	JenisDayaTarik  []*string          `json:"jenis_daya_tarik" bson:"jenis_daya_tarik,omitempty"`
	JenisAktifitas  []*string          `json:"jenis_aktifitas" bson:"jenis_aktifitas,omitempty"`
	Alamat          string             `json:"alamat" bson:"alamat,omitempty"`
	Koordinat       *Koordinat         `json:"koordinat" bson:"koordinat,omitempty"`
	MediaSosial     *MediaSosial       `json:"media_sosial" bson:"media_sosial,omitempty"`
	Tags            []*string          `json:"tags" bson:"tags,omitempty"`
	IsDelete        bool               `json:"is_delete" bson:"is_delete"`
	IsActive        bool               `json:"is_active" bson:"is_active"`
	Log             *Log               `json:"log" bson:"log,omitempty"`
}

type WisataRepository interface {
	Add(ctx context.Context, d *Wisata) (primitive.ObjectID, r.Ex)
	FetchAll(ctx context.Context) ([]*Wisata, r.Ex)
	FetchByID(ctx context.Context, ID primitive.ObjectID) (*Wisata, r.Ex)
	UpdateByID(ctx context.Context, ID primitive.ObjectID, d *Wisata) r.Ex
	DeleteByID(ctx context.Context, ID primitive.ObjectID) r.Ex
	Active(ctx context.Context, ID primitive.ObjectID) r.Ex
	FetchBySlug(ctx context.Context, slug string) (*Wisata, r.Ex)
}
