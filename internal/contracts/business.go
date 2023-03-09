package contracts

import (
	"62teknologi-senior-backend-test-muhammad_risky_nugraha/internal/domain"
	"context"
)

type BusinessUseCase interface {
	GetBusinessData(ctx context.Context) (*[]domain.Business, error)
}