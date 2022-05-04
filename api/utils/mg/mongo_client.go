package mg

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoClient interface {
	Client() *mongo.Client
	TranscationSession() (TranscationSession, error)
}

type mongoClient struct {
	client *mongo.Client
}

func NewMongoClient(c *mongo.Client) MongoClient {
	return &mongoClient{
		client: c,
	}
}

func (m *mongoClient) Client() *mongo.Client {
	return m.client
}

func (m *mongoClient) TranscationSession() (TranscationSession, error) {
	//txOptions := options.Session().SetDefaultWriteConcern(wc).SetDefaultReadConcern(rc)
	session, errIn := m.Client().StartSession()
	if errIn != nil {
		return nil, errIn
	}
	return NewTranscationSession(session), nil
}
