package mg

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type TranscationSessionMock struct {
	mock.Mock
}

func (t *TranscationSessionMock) CommitTransaction(tc TranscationContext) error {
	arguments := t.Mock.Called(tc)
	return arguments.Error(0)
}

func (t *TranscationSessionMock) AbortTransaction(ctx context.Context) error {
	arguments := t.Mock.Called(ctx)
	return arguments.Error(0)
}

func (t *TranscationSessionMock) StartTransaction(ctx context.Context, fn func(tc TranscationContext) error) error {
	return fn(&TranscationContextMock{})
}

func (t *TranscationSessionMock) EndSession(ctx context.Context) {
	//arguments := t.Mock.Called(ctx)
}
