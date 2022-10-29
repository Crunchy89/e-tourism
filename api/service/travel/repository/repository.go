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

type travelRepository struct {
	db   *mongo.Database
	coll *mongo.Collection
}

func NewTravelRepository(db *mongo.Database) domain.TravelRepository {
	return &travelRepository{
		db:   db,
		coll: db.Collection("travel"),
	}
}

func (p *travelRepository) consumeCursor(ctx context.Context, cur *mongo.Cursor) ([]*domain.Travel, r.Ex) {
	defer cur.Close(ctx)

	res := make([]*domain.Travel, 0)
	for cur.Next(ctx) {
		h := new(domain.Travel)
		if err := cur.Decode(h); err != nil {
			return res, r.NewErrorMongo(p.coll.Name(), err)
		}
		res = append(res, h)
	}
	return res, nil
}

func (p *travelRepository) insert(ctx context.Context, d interface{}) (primitive.ObjectID, r.Ex) {
	res, err := p.coll.InsertOne(ctx, d)
	if err != nil {
		return primitive.NilObjectID, r.NewErrorMongo(p.coll.Name(), err)
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

func (p *travelRepository) findOne(ctx context.Context, d interface{}, opts ...*options.FindOneOptions) (*domain.Travel, r.Ex) {
	res := new(domain.Travel)
	if err := p.coll.FindOne(ctx, d).Decode(res); err != nil {
		return nil, r.NewErrorMongo(p.coll.Name(), err)
	}

	return res, nil
}

func (p *travelRepository) updatedOne(ctx context.Context, f interface{}, u interface{}, opts ...*options.UpdateOptions) r.Ex {
	if _, err := p.coll.UpdateOne(ctx, f, u, opts...); err != nil {
		return r.NewErrorMongo(p.coll.Name(), err)
	}
	return nil
}

func (p *travelRepository) findForUpdate(ctx context.Context, f interface{}, opts ...*options.FindOneAndUpdateOptions) (*domain.Travel, r.Ex) {
	updated := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "_lock_tx", Value: primitive.NewObjectID()},
		}},
	}
	res := new(domain.Travel)
	if err := p.coll.FindOneAndUpdate(ctx, f, updated, opts...).Decode(res); err != nil {
		return res, r.NewErrorMongo(p.coll.Name(), err)
	}

	return res, nil
}

func (p *travelRepository) find(ctx context.Context, d interface{}, opts ...*options.FindOptions) ([]*domain.Travel, r.Ex) {
	cur, err := p.coll.Find(ctx, d, opts...)
	if err != nil {
		return []*domain.Travel{}, r.NewErrorMongo(p.coll.Name(), err)
	}

	return p.consumeCursor(ctx, cur)
}

// func add data
func (p *travelRepository) Add(ctx context.Context, d *domain.Travel) (primitive.ObjectID, r.Ex) {
	return p.insert(ctx, d)
}

// func fetch all
func (p *travelRepository) FetchAll(ctx context.Context) ([]*domain.Travel, r.Ex) {
	return p.find(ctx, bson.M{"is_delete": false, "is_active": true})
}

// func fetch by id
func (p *travelRepository) FetchByID(ctx context.Context, ID primitive.ObjectID) (*domain.Travel, r.Ex) {
	return p.findOne(ctx, bson.M{"_id": ID, "is_delete": false, "is_active": true})
}

// funct fetch by slug
func (p *travelRepository) FetchBySlug(ctx context.Context, slug string) (*domain.Travel, r.Ex) {
	return p.findOne(ctx, bson.M{"slug": slug, "is_delete": false, "is_active": true})
}

// func update by id
func (p *travelRepository) UpdateByID(ctx context.Context, ID primitive.ObjectID, d *domain.Travel) r.Ex {
	return p.updatedOne(ctx, bson.M{"_id": ID}, d)
}

// function delete set status is_delete to true
func (p *travelRepository) DeleteByID(ctx context.Context, ID primitive.ObjectID) r.Ex {
	return p.updatedOne(ctx, bson.M{"_id": ID}, bson.M{"$set": bson.M{"is_delete": true}})
}

// function Active set status is_active to true
func (p *travelRepository) ActiveByID(ctx context.Context, ID primitive.ObjectID) r.Ex {
	return p.updatedOne(ctx, bson.M{"_id": ID}, bson.M{"$set": bson.M{"is_active": true}})
}
