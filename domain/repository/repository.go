package repository

import (
	"context"
	"yourwishes/domain/entity"
)

type SaveWishesRepo interface {
	SaveWishes(ctx context.Context, obj *entity.Wishes) error
}

type FindAllWishesRepo interface {
	FindAllWishes(ctx context.Context) ([]*entity.Wishes, error)
}
