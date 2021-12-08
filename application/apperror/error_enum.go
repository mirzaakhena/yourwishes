package apperror

const (
	FailUnmarshalResponseBodyError ErrorType = "ER1000 Fail to unmarshal response body"   // used by controller
	EntityNotFound                 ErrorType = "ER1001 Entity %s with id %s is not found" // used by injected repo in interactor
	UnrecognizedEnum               ErrorType = "ER1002 %s is not recognized %s enum"      // used by enum
	DatabaseNotFoundInContextError ErrorType = "ER1003 Database is not found in context"  // used by repoimpl
)
const WishesMessageLengthHasExceeded ErrorType = "ER1000 wishes message length has exceeded" //
