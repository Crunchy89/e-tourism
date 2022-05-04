package mg

import (
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoClientMock struct {
	mock.Mock
}

func (m *MongoClientMock) Client() *mongo.Client {
	return nil
}

func (m *MongoClientMock) TranscationSession() (TranscationSession, error) {
	arguments := m.Mock.Called()
	if arguments.Get(0) != nil {
		return arguments.Get(0).(TranscationSession), nil
	} else {
		return nil, arguments.Error(1)
	}
}
