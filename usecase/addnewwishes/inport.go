package addnewwishes

import (
	"context"
	"time"
)

// Inport of Usecase
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase
type InportRequest struct {
	Message string
	Now     time.Time
}

// InportResponse is response payload after running the usecase
type InportResponse struct {
}
