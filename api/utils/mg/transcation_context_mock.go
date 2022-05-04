package mg

import (
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

type TranscationContextMock struct {
	mock.Mock
}

func (t *TranscationContextMock) SessionContext() mongo.SessionContext {
	return nil
}
