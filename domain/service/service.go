package service

import "context"

type GenerateIDService interface {
	GenerateID(ctx context.Context) string
}
