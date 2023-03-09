package usecase

import (
	"62teknologi-senior-backend-test-muhammad_risky_nugraha/internal/domain"

	"github.com/stretchr/testify/mock"

	// "gitlab.com/hmcorp/wallet-common/internal/utils/filtering"
	"github.com/gofiber/fiber/v2"
)


type BusinessUsecaseMock struct {
	mock.Mock
}

func (r *BusinessUsecaseMock) GetAllData(ctx *fiber.Ctx) ([]*domain.Business, error) {
	ret := r.Called(ctx)

	var res []*domain.Business
	if rf, ok := ret.Get(0).(func(*fiber.Ctx) []*domain.Business); ok {
		res = rf(ctx)
	} else {
		res = ret.Get(0).([]*domain.Business)
	}

	var err error

	if rf, ok := ret.Get(1).(func(*fiber.Ctx) error); ok {
		err = rf(ctx)
	} else {
		err = ret.Error(1)
	}

	return res, err

	// if rf, ok := ret.Get(0)
}