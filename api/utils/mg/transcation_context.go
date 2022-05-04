package mg

import "go.mongodb.org/mongo-driver/mongo"

type TranscationContext interface {
	SessionContext() mongo.SessionContext
}

type transcationContext struct {
	sc mongo.SessionContext
}

func (t *transcationContext) SessionContext() mongo.SessionContext {
	return t.sc
}

func NewTranscationContext(sc mongo.SessionContext) TranscationContext {
	return &transcationContext{
		sc: sc,
	}
}
