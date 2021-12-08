package displayallwishes

import (
	"context"
	"yourwishes/domain/service"
)

//go:generate mockery --name Outport -output mocks/

type displayAllWishesInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &displayAllWishesInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *displayAllWishesInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	err := service.WithoutTransaction(ctx, r.outport, func(ctx context.Context) error {
		wishesObjs, err := r.outport.FindAllWishes(ctx)
		if err != nil {
			return err
		}

		for _, obj := range wishesObjs {
			res.ListOfWishes = append(res.ListOfWishes, Wishes{
				ID:      obj.ID.String(),
				Message: obj.Message,
				Date:    obj.Now,
			})
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
