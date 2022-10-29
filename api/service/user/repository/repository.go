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
		{Key: "$set", Value: bson.D{
			{Key: "_lock_tx", Value: primitive.NewObjectID()},
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

// func add data
func (p *userRepository) Add(ctx context.Context, d *domain.User) (primitive.ObjectID, r.Ex) {
	return p.insert(ctx, d)
}

// func fetch all
func (p *userRepository) FetchAll(ctx context.Context) ([]*domain.User, r.Ex) {
	return p.find(ctx, bson.M{"is_delete": false, "is_active": true})
}

// func fetch by id
func (p *userRepository) FetchByID(ctx context.Context, ID primitive.ObjectID) (*domain.User, r.Ex) {
	return p.findOne(ctx, bson.M{"_id": ID, "is_delete": false, "is_active": true})
}

// fetch by username
func (p *userRepository) FetchByUsername(ctx context.Context, username string) (*domain.User, r.Ex) {
	return p.findOne(ctx, bson.M{"username": username})
}

// fetch by email
func (p *userRepository) FetchByEmail(ctx context.Context, email string) (*domain.User, r.Ex) {
	return p.findOne(ctx, bson.M{"email": email, "is_delete": false, "is_active": true})
}

// func update by id
func (p *userRepository) UpdateByID(ctx context.Context, ID primitive.ObjectID, d *domain.User) r.Ex {
	return p.updatedOne(ctx, bson.M{"_id": ID}, d)
}

// func update by username
func (p *userRepository) UpdateByUsername(ctx context.Context, username string, d *domain.User) r.Ex {
	return p.updatedOne(ctx, bson.M{"username": username}, d)
}

// func update by email
func (p *userRepository) UpdateByEmail(ctx context.Context, email string, d *domain.User) r.Ex {
	return p.updatedOne(ctx, bson.M{"email": email}, d)
}

// function delete set status is_delete to true
func (p *userRepository) DeleteByID(ctx context.Context, ID primitive.ObjectID) r.Ex {
	return p.updatedOne(ctx, bson.M{"_id": ID}, bson.M{"$set": bson.M{"is_delete": true}})
}

// function Active set status is_active to true
func (p *userRepository) ActiveByID(ctx context.Context, ID primitive.ObjectID) r.Ex {
	return p.updatedOne(ctx, bson.M{"_id": ID}, bson.M{"$set": bson.M{"is_active": true}})
}

// function pagination
func (p *userRepository) Pagination(ctx context.Context, page int64, limit int64) ([]*domain.User, r.Ex) {
	skip := (page - 1) * limit
	return p.find(ctx, bson.M{"is_delete": false}, &options.FindOptions{
		Sort:  bson.M{"created_at": -1},
		Skip:  &skip,
		Limit: &limit,
	})
}

// fucnction search with pagination
func (p *userRepository) Search(ctx context.Context, page int64, limit int64, search string) ([]*domain.User, r.Ex) {
	skip := (page - 1) * limit
	return p.find(ctx, bson.M{"is_delete": false, "$or": bson.A{bson.M{"username": bson.M{"$regex": search}}, bson.M{"email": bson.M{"$regex": search}}}}, &options.FindOptions{
		Sort:  bson.M{"created_at": -1},
		Skip:  &skip,
		Limit: &limit,
	})
}

// function count all where is_delete is false
func (p *userRepository) CountAll(ctx context.Context) (int64, r.Ex) {
	count, err := p.coll.CountDocuments(ctx, bson.M{"is_delete": false})
	if err != nil {
		return 0, r.NewErrorMongo(p.coll.Name(), err)
	}
	return count, nil
}
