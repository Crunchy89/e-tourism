package repository

import (
	"context"

	"api/domain"
	"api/utils/r"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type penginapanRepository struct {
	db   *mongo.Database
	coll *mongo.Collection
}

func NewPenginapanRepository(db *mongo.Database) domain.PenginapanRepository {
	return &penginapanRepository{
		db:   db,
		coll: db.Collection("penginapan"),
	}
}

func (p *penginapanRepository) consumeCursor(ctx context.Context, cur *mongo.Cursor) ([]*domain.Penginapan, r.Ex) {
	defer cur.Close(ctx)

	res := make([]*domain.Penginapan, 0)
	for cur.Next(ctx) {
		h := new(domain.Penginapan)
		if err := cur.Decode(h); err != nil {
			return res, r.NewErrorMongo(p.coll.Name(), err)
		}
		res = append(res, h)
	}
	return res, nil
}

// func insert
func (p *penginapanRepository) insert(ctx context.Context, d interface{}) (primitive.ObjectID, r.Ex) {
	res, err := p.coll.InsertOne(ctx, d)
	if err != nil {
		return primitive.NilObjectID, r.NewErrorMongo(p.coll.Name(), err)
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

// func insert many
func (p *penginapanRepository) insertMany(ctx context.Context, d []interface{}) ([]primitive.ObjectID, r.Ex) {
	res, err := p.coll.InsertMany(ctx, d)
	if err != nil {
		return []primitive.ObjectID{}, r.NewErrorMongo(p.coll.Name(), err)
	}
	// convert to slice of primitive.ObjectID
	IDs := make([]primitive.ObjectID, 0)
	for _, id := range res.InsertedIDs {
		IDs = append(IDs, id.(primitive.ObjectID))
	}
	return IDs, nil
}

// func findOne
func (p *penginapanRepository) findOne(ctx context.Context, d interface{}, opts ...*options.FindOneOptions) (*domain.Penginapan, r.Ex) {
	res := new(domain.Penginapan)
	if err := p.coll.FindOne(ctx, d).Decode(res); err != nil {
		return nil, r.NewErrorMongo(p.coll.Name(), err)
	}

	return res, nil
}

// func find many
func (p *penginapanRepository) find(ctx context.Context, d interface{}, opts ...*options.FindOptions) ([]*domain.Penginapan, r.Ex) {
	cur, err := p.coll.Find(ctx, d, opts...)
	if err != nil {
		return []*domain.Penginapan{}, r.NewErrorMongo(p.coll.Name(), err)
	}

	return p.consumeCursor(ctx, cur)
}

// func update one
func (p *penginapanRepository) updatedOne(ctx context.Context, f interface{}, u interface{}, opts ...*options.UpdateOptions) r.Ex {
	if _, err := p.coll.UpdateOne(ctx, f, u, opts...); err != nil {
		return r.NewErrorMongo(p.coll.Name(), err)
	}
	return nil
}

// func update many
func (p *penginapanRepository) UpdateMany(ctx context.Context, f interface{}, u interface{}, opts ...*options.UpdateOptions) r.Ex {
	if _, err := p.coll.UpdateMany(ctx, f, u, opts...); err != nil {
		return r.NewErrorMongo(p.coll.Name(), err)
	}
	return nil
}

// func find data to update
func (p *penginapanRepository) findForUpdate(ctx context.Context, f interface{}, opts ...*options.FindOneAndUpdateOptions) (*domain.Penginapan, r.Ex) {
	updated := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "_lock_tx", Value: primitive.NewObjectID()},
		}},
	}
	res := new(domain.Penginapan)
	if err := p.coll.FindOneAndUpdate(ctx, f, updated, opts...).Decode(res); err != nil {
		return res, r.NewErrorMongo(p.coll.Name(), err)
	}

	return res, nil
}

// func add new data
func (p *penginapanRepository) Add(ctx context.Context, d *domain.Penginapan) (primitive.ObjectID, r.Ex) {
	return p.insert(ctx, d)
}

// func delete by id
func (p *penginapanRepository) Delete(ctx context.Context, ID primitive.ObjectID) r.Ex {
	if _, err := p.coll.DeleteOne(ctx, bson.M{"_id": ID}); err != nil {
		return r.NewErrorMongo(p.coll.Name(), err)
	}
	return nil
}

// func fetch all
func (p *penginapanRepository) FetchAll(ctx context.Context) ([]*domain.Penginapan, r.Ex) {
	return p.find(ctx, bson.M{"is_deleted": false})
}

// func fetch by id
func (p *penginapanRepository) FetchByID(ctx context.Context, ID primitive.ObjectID) (*domain.Penginapan, r.Ex) {
	filter := bson.M{"_id": ID, "is_deleted": false}
	return p.findOne(ctx, filter)
}

// func fetch by slug
func (p *penginapanRepository) FetchBySlug(ctx context.Context, slug string) (*domain.Penginapan, r.Ex) {
	return p.findOne(ctx, bson.M{"slug": slug, "is_deleted": false})
}

// funct update
func (p *penginapanRepository) UpdateByID(ctx context.Context, ID primitive.ObjectID, d *domain.Penginapan) r.Ex {
	filter := bson.M{"_id": ID}
	return p.updatedOne(ctx, filter, d)
}

// funct delete
func (p *penginapanRepository) DeleteByID(ctx context.Context, ID primitive.ObjectID) r.Ex {
	filter := bson.M{"_id": ID}
	return p.updatedOne(ctx, filter, bson.M{"is_deleted": true})
}

// func active
func (p *penginapanRepository) Active(ctx context.Context, ID primitive.ObjectID) r.Ex {
	filter := bson.M{"_id": ID}
	return p.updatedOne(ctx, filter, bson.M{"is_active": true})
}

// function pagination
func (p *penginapanRepository) Pagination(ctx context.Context, page int64, limit int64) ([]*domain.Penginapan, r.Ex) {
	skip := (page - 1) * limit
	return p.find(ctx, bson.M{"is_delete": false}, &options.FindOptions{
		Sort:  bson.M{"created_at": -1},
		Skip:  &skip,
		Limit: &limit,
	})
}

// fucnction search with pagination
func (p *penginapanRepository) Search(ctx context.Context, page int64, limit int64, search string) ([]*domain.Penginapan, r.Ex) {
	skip := (page - 1) * limit
	return p.find(ctx, bson.M{"is_delete": false, "$or": bson.A{bson.M{"penginapanname": bson.M{"$regex": search}}, bson.M{"email": bson.M{"$regex": search}}}}, &options.FindOptions{
		Sort:  bson.M{"created_at": -1},
		Skip:  &skip,
		Limit: &limit,
	})
}

// function count all where is_delete is false
func (p *penginapanRepository) CountAll(ctx context.Context) (int64, r.Ex) {
	count, err := p.coll.CountDocuments(ctx, bson.M{"is_delete": false})
	if err != nil {
		return 0, r.NewErrorMongo(p.coll.Name(), err)
	}
	return count, nil
}
