package addnewwishes

import (
	"yourwishes/domain/repository"
	"yourwishes/domain/service"
)

// Outport of usecase
type Outport interface {
	repository.SaveWishesRepo
	service.GenerateIDService
	repository.WithTransactionDB
}
