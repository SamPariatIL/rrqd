package utils

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionManager interface {
	StartTransaction(ctx context.Context) (mongo.SessionContext, error)
	CommitTransaction(ctx mongo.SessionContext)
	AbortTransaction(ctx mongo.SessionContext)
}

type transactionManager struct {
	db *mongo.Database
}

func NewTransactionManager(db *mongo.Database) TransactionManager {
	return &transactionManager{db: db}
}

func (tm *transactionManager) StartTransaction(ctx context.Context) (mongo.SessionContext, error) {
	session, err := tm.db.Client().StartSession()
	if err != nil {
		return nil, err
	}

	err = session.StartTransaction()
	if err != nil {
		session.EndSession(ctx)
		return nil, err
	}

	return mongo.NewSessionContext(ctx, session), nil
}

func (tm *transactionManager) CommitTransaction(ctx mongo.SessionContext) {
	_ = ctx.CommitTransaction(ctx)
	ctx.EndSession(ctx)
}

func (tm *transactionManager) AbortTransaction(ctx mongo.SessionContext) {
	_ = ctx.AbortTransaction(ctx)
	ctx.EndSession(ctx)
}
