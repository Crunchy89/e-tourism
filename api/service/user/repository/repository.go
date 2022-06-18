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

type userRepository struct {
	db   *mongo.Database
	coll *mongo.Collection
}

func NewUserRepository(db *mongo.Database) domain.UserRepository {
	return &userRepository{
		db:   db,
		coll: db.Collection("user"),
	}
}

func (p *userRepository) consumeCursor(ctx context.Context, cur *mongo.Cursor) ([]*domain.User, r.Ex) {
	defer cur.Close(ctx)

	res := make([]*domain.User, 0)
	for cur.Next(ctx) {
		h := new(domain.User)
		if err := cur.Decode(h); err != nil {
			return res, r.NewErrorMongo(p.coll.Name(), err)
		}
		res = append(res, h)
	}
	return res, nil
}

func (p *userRepository) insert(ctx context.Context, d interface{}) (primitive.ObjectID, r.Ex) {
	res, err := p.coll.InsertOne(ctx, d)
	if err != nil {
		return primitive.NilObjectID, r.NewErrorMongo(p.coll.Name(), err)
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

func (p *userRepository) findOne(ctx context.Context, d interface{}, opts ...*options.FindOneOptions) (*domain.User, r.Ex) {
	res := new(domain.User)
	if err := p.coll.FindOne(ctx, d).Decode(res); err != nil {
		return nil, r.NewErrorMongo(p.coll.Name(), err)
	}

	return res, nil
}

func (p *userRepository) updatedOne(ctx context.Context, f interface{}, u interface{}, opts ...*options.UpdateOptions) r.Ex {
	if _, err := p.coll.UpdateOne(ctx, f, u, opts...); err != nil {
		return r.NewErrorMongo(p.coll.Name(), err)
	}
	return nil
}

func (p *userRepository) findForUpdate(ctx context.Context, f interface{}, opts ...*options.FindOneAndUpdateOptions) (*domain.User, r.Ex) {
	updated := bson.D{
		{"$set", bson.D{
			{"_lock_tx", primitive.NewObjectID()},
		}},
	}
	res := new(domain.User)
	if err := p.coll.FindOneAndUpdate(ctx, f, updated, opts...).Decode(res); err != nil {
		return res, r.NewErrorMongo(p.coll.Name(), err)
	}

	return res, nil
}

func (p *userRepository) find(ctx context.Context, d interface{}, opts ...*options.FindOptions) ([]*domain.User, r.Ex) {
	cur, err := p.coll.Find(ctx, d, opts...)
	if err != nil {
		return []*domain.User{}, r.NewErrorMongo(p.coll.Name(), err)
	}

	return p.consumeCursor(ctx, cur)
}

func (p *userRepository) Add(ctx context.Context, d *domain.User) (primitive.ObjectID, r.Ex) {
	return p.insert(ctx, d)
}

func (p *userRepository) FetchAll(ctx context.Context) ([]*domain.User, r.Ex) {
	return p.find(ctx, bson.M{})
}

func (p *userRepository) Fetch(ctx context.Context, ID primitive.ObjectID) (*domain.User, r.Ex) {
	return p.findOne(ctx, bson.M{"_id": ID})
}

func (p *userRepository) FetchByUsername(ctx context.Context, username string) (*domain.User, r.Ex) {
	return p.findOne(ctx, bson.M{"username": username})
}
