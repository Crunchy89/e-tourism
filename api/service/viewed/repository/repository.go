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

type viewedRepository struct {
	db   *mongo.Database
	coll *mongo.Collection
}

func NewViewedRepository(db *mongo.Database) domain.ViewedRepository {
	return &viewedRepository{
		db:   db,
		coll: db.Collection("viewed"),
	}
}

func (p *viewedRepository) consumeCursor(ctx context.Context, cur *mongo.Cursor) ([]*domain.Viewed, r.Ex) {
	defer cur.Close(ctx)

	res := make([]*domain.Viewed, 0)
	for cur.Next(ctx) {
		h := new(domain.Viewed)
		if err := cur.Decode(h); err != nil {
			return res, r.NewErrorMongo(p.coll.Name(), err)
		}
		res = append(res, h)
	}
	return res, nil
}

// func insert
func (p *viewedRepository) insert(ctx context.Context, d interface{}) (primitive.ObjectID, r.Ex) {
	res, err := p.coll.InsertOne(ctx, d)
	if err != nil {
		return primitive.NilObjectID, r.NewErrorMongo(p.coll.Name(), err)
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

// func insert many
func (p *viewedRepository) insertMany(ctx context.Context, d []interface{}) ([]primitive.ObjectID, r.Ex) {
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
func (p *viewedRepository) findOne(ctx context.Context, d interface{}, opts ...*options.FindOneOptions) (*domain.Viewed, r.Ex) {
	res := new(domain.Viewed)
	if err := p.coll.FindOne(ctx, d).Decode(res); err != nil {
		return nil, r.NewErrorMongo(p.coll.Name(), err)
	}

	return res, nil
}

// func find many
func (p *viewedRepository) find(ctx context.Context, d interface{}, opts ...*options.FindOptions) ([]*domain.Viewed, r.Ex) {
	cur, err := p.coll.Find(ctx, d, opts...)
	if err != nil {
		return []*domain.Viewed{}, r.NewErrorMongo(p.coll.Name(), err)
	}

	return p.consumeCursor(ctx, cur)
}

// func update one
func (p *viewedRepository) updatedOne(ctx context.Context, f interface{}, u interface{}, opts ...*options.UpdateOptions) r.Ex {
	if _, err := p.coll.UpdateOne(ctx, f, u, opts...); err != nil {
		return r.NewErrorMongo(p.coll.Name(), err)
	}
	return nil
}

// func update many
func (p *viewedRepository) UpdateMany(ctx context.Context, f interface{}, u interface{}, opts ...*options.UpdateOptions) r.Ex {
	if _, err := p.coll.UpdateMany(ctx, f, u, opts...); err != nil {
		return r.NewErrorMongo(p.coll.Name(), err)
	}
	return nil
}

// func find data to update
func (p *viewedRepository) findForUpdate(ctx context.Context, f interface{}, opts ...*options.FindOneAndUpdateOptions) (*domain.Viewed, r.Ex) {
	updated := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "_lock_tx", Value: primitive.NewObjectID()},
		}},
	}
	res := new(domain.Viewed)
	if err := p.coll.FindOneAndUpdate(ctx, f, updated, opts...).Decode(res); err != nil {
		return res, r.NewErrorMongo(p.coll.Name(), err)
	}

	return res, nil
}

// func add new data
func (p *viewedRepository) Add(ctx context.Context, d *domain.Viewed) (primitive.ObjectID, r.Ex) {
	return p.insert(ctx, d)
}

// func delete by id
func (p *viewedRepository) Delete(ctx context.Context, ID primitive.ObjectID) r.Ex {
	if _, err := p.coll.DeleteOne(ctx, bson.M{"_id": ID}); err != nil {
		return r.NewErrorMongo(p.coll.Name(), err)
	}
	return nil
}

// func fetch all
func (p *viewedRepository) FetchAll(ctx context.Context) ([]*domain.Viewed, r.Ex) {
	return p.find(ctx, bson.M{"is_deleted": false})
}

// func fetch by id
func (p *viewedRepository) FetchByID(ctx context.Context, ID primitive.ObjectID) (*domain.Viewed, r.Ex) {
	filter := bson.M{"_id": ID, "is_deleted": false}
	return p.findOne(ctx, filter)
}

// funct update
func (p *viewedRepository) UpdateByID(ctx context.Context, ID primitive.ObjectID, d *domain.Viewed) r.Ex {
	filter := bson.M{"_id": ID}
	return p.updatedOne(ctx, filter, d)
}

// funct delete
func (p *viewedRepository) DeleteByID(ctx context.Context, ID primitive.ObjectID) r.Ex {
	filter := bson.M{"_id": ID}
	return p.updatedOne(ctx, filter, bson.M{"is_deleted": true})
}
