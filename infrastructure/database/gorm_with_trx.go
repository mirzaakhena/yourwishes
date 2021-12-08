package database

import (
	"context"
	"gorm.io/gorm"
	"yourwishes/infrastructure/log"
)

type GormWithTransactionImpl struct {
	db *gorm.DB
}

func NewGormWithTransactionImpl(db *gorm.DB) *GormWithTransactionImpl {
	return &GormWithTransactionImpl{
		db: db,
	}
}

func (r *GormWithTransactionImpl) BeginTransaction(ctx context.Context) (context.Context, error) {
	dbTrx := r.db.Begin()
	trxCtx := context.WithValue(ctx, ContextDBValue, dbTrx)
	return trxCtx, nil
}

func (r *GormWithTransactionImpl) CommitTransaction(ctx context.Context) error {
	log.Info(ctx, "Commit Transaction")
	db, err := ExtractDB(ctx)
	if err != nil {
		return err
	}
	return db.Commit().Error
}

func (r *GormWithTransactionImpl) RollbackTransaction(ctx context.Context) error {
	log.Info(ctx, "Rollback Transaction")

	db, err := ExtractDB(ctx)
	if err != nil {
		return err
	}

	return db.Rollback().Error
}
