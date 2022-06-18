package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Log struct {
	Created primitive.DateTime `json:"created" bson:"created,omitempty"`
	Updated primitive.DateTime `json:"updated" bson:"updated,omitempty"`
}

type Koordinat struct {
	Longitude float64 `json:"longitude" bson:"longitude,omitempty"`
	Latitude  float64 `json:"latitude" bson:"latitude,omitempty"`
}

type Photo struct {
	Path string `json:"path" bson:"path,omitempty"`
}

type Thumbnails struct {
	Path string `json:"path" bson:"path,omitempty"`
}

type JenisWisata struct {
	Jenis string `json:"jenis" bson:"jenis,omitempty"`
}

type JenisAktifitas struct {
	Jenis string `json:"jenis" bson:"jenis,omitempty"`
}

type Tags struct {
	Tags string `json:"tags" bson:"tags,omitempty"`
}
type MediaSosial struct {
	Facebook  string `json:"facebook" bson:"facebook,omitempty"`
	Instagram string `json:"instagram" bson:"instagram,omitempty"`
	Twitter   string `json:"twitter" bson:"twitter,omitempty"`
	Youtube   string `json:"youtube" bson:"youtube,omitempty"`
}
