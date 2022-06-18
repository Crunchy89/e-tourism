package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Wisata struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	NamaObjekWisata string             `json:"nama_objek_wisata" bson:"nama_objek_wisata,omitempty"`
	Slug            string             `json:"slug" bson:"slug,omitempty"`
	PotensiWisata   string             `json:"potensi_wisata" bson:"potensi_wisata,omitempty"`
	JenisDayaTarik  []*JenisWisata     `json:"jenis_daya_tarik" bson:"jenis_daya_tarik,omitempty"`
	JenisAktifitas  []*JenisAktifitas  `json:"jenis_aktifitas" bson:"jenis_aktifitas,omitempty"`
	Alamat          string             `json:"alamat" bson:"alamat,omitempty"`
	Koordinat       *Koordinat         `json:"koordinat" bson:"koordinat,omitempty"`
	Photo           []*Photo           `json:"photo" bson:"photo,omitempty"`
	Thumbnails      []*Thumbnails      `json:"thumbnails" bson:"thumbnails,omitempty"`
	MediaSosial     *MediaSosial       `json:"media_sosial" bson:"media_sosial,omitempty"`
	Tags            []*Tags            `json:"tags" bson:"tags,omitempty"`
	IsDelete        bool               `json:"is_delete" bson:"is_delete"`
	IsActive        bool               `json:"is_active" bson:"is_active"`
	Log             *Log               `json:"log" bson:"log,omitempty"`
}
