package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Kamar struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	KamarID     primitive.ObjectID `json:"kamar_id" bson:"kamar_id,omitempty"`
	JumlahKamar uint16             `json:"jumlah_kamar" bson:"jumlah_kamar,omitempty"`
	Harga       uint64             `json:"harga" bson:"harga,omitempty"`
	Diskon      uint64             `json:"diskon" bson:"diskon,omitempty"`
	Photo       []*Photo           `json:"photo" bson:"photo,omitempty"`
	Thumbnails  []*Thumbnails      `json:"thumbnails" bson:"thumbnails,omitempty"`
	IsDelete    bool               `json:"is_delete" bson:"is_delete"`
	IsActive    bool               `json:"is_active" bson:"is_active"`
	Log         *Log               `json:"log" bson:"log,omitempty"`
}
