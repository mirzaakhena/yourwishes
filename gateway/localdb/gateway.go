package localdb

import (
	"context"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"yourwishes/infrastructure/database"

	"yourwishes/domain/entity"
	"yourwishes/infrastructure/log"
)

type gateway struct {
	*database.GormWithTransactionImpl
	*database.GormWithoutTransactionImpl
}

// NewGateway ...
func NewGateway() *gateway {

	db := database.NewSQLiteDefault()

	err := db.AutoMigrate(&entity.Wishes{})
	if err != nil {
		panic(err)
	}

	return &gateway{
		GormWithTransactionImpl: database.NewGormWithTransactionImpl(db),
		GormWithoutTransactionImpl: database.NewGormWithoutTransactionImpl(db),
	}
}

func (r *gateway) SaveWishes(ctx context.Context, obj *entity.Wishes) error {
	log.Info(ctx, "called")

	db, err := database.ExtractDB(ctx)
	if err != nil {
		return err
	}

	err = db.Save(obj).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *gateway) GenerateID(ctx context.Context) string {
	log.Info(ctx, "called")

	id, err := gonanoid.Generate("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890", 4)
	if err != nil {
		return "abcd"
	}

	return id
}

func (r *gateway) FindAllWishes(ctx context.Context) ([]*entity.Wishes, error) {
	db, err := database.ExtractDB(ctx)
	if err != nil {
		return nil, err
	}

	objs := make([]*entity.Wishes, 0)

	err = db.Find(&objs).Error
	if err != nil {
		return nil, err
	}

	return objs, nil
}
