package entity

import (
	"time"
	"yourwishes/application/apperror"
	"yourwishes/domain/vo"
)

type Wishes struct {
	ID      vo.WishesID
	Message string
	Now     time.Time
}

type WishesRequest struct {
	RandomID string
	Message  string
	Now      time.Time
}

func NewWishes(req WishesRequest) (*Wishes, error) {

	var obj Wishes

	id, err := vo.NewWishesID(req.RandomID)
	if err != nil {
		return nil, err
	}

	obj.Message = req.Message
	obj.Now = req.Now
	obj.ID = id

	err = obj.Validate()
	if err != nil {
		return nil, err
	}

	return &obj, nil
}

func (r *Wishes) Validate() error {

	if len(r.Message) > 60 {
		return apperror.WishesMessageLengthHasExceeded
	}

	return nil
}
