package displayallwishes

import "yourwishes/domain/repository"

// Outport of usecase
type Outport interface {
	repository.FindAllWishesRepo
	repository.WithoutTransactionDB
}
