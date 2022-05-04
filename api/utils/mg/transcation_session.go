package mg

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type TranscationSession interface {
	CommitTransaction(tc TranscationContext) error
	AbortTransaction(context.Context) error
	StartTransaction(ctx context.Context, fn func(tc TranscationContext) error) error
	EndSession(ctx context.Context)
}

type transcationSession struct {
	session mongo.Session
}

func NewTranscationSession(s mongo.Session) TranscationSession {
	return &transcationSession{
		session: s,
	}
}

func (t *transcationSession) CommitTransaction(tc TranscationContext) error {
	//return t.session.CommitTransaction(tc.SessionContext())
	return nil
}

func (t *transcationSession) AbortTransaction(ctx context.Context) error {
	//return t.session.AbortTransaction(ctx)
	return nil
}

func (t *transcationSession) StartTransaction(ctx context.Context, fn func(tc TranscationContext) error) error {
	wc := writeconcern.New(writeconcern.WMajority(), writeconcern.J(true))
	rc := readconcern.Snapshot()
	_, err := t.session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		return nil, fn(NewTranscationContext(sessCtx))
	}, options.Transaction().SetReadConcern(rc).SetWriteConcern(wc))
	return err
	// return mongo.WithSession(ctx, t.session, func(sc mongo.SessionContext) error {
	// 	if err := t.session.StartTransaction(); err != nil {
	// 		return err
	// 	}

	// 	return fn(NewTranscationContext(sc))
	// })
}

func (t *transcationSession) EndSession(ctx context.Context) {
	t.session.EndSession(ctx)
}
