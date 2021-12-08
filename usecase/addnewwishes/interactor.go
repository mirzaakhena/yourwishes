package addnewwishes

import (
	"context"
	"yourwishes/domain/entity"
	"yourwishes/domain/service"
)

//go:generate mockery --name Outport -output mocks/

type addNewWishesInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &addNewWishesInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *addNewWishesInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	err := service.WithTransaction(ctx, r.outport, func(ctx context.Context) error {

		randomID := r.outport.GenerateID(ctx)

		wishesObj, err := entity.NewWishes(entity.WishesRequest{
			RandomID: randomID,
			Message:  req.Message,
			Now:      req.Now,
		})
		if err != nil {
			return err
		}

		err = r.outport.SaveWishes(ctx, wishesObj)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
