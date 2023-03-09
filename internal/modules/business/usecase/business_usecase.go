package usecase

import (
	"context"
	"fmt"

	"62teknologi-senior-backend-test-muhammad_risky_nugraha/internal/domain"
	businessRepo "62teknologi-senior-backend-test-muhammad_risky_nugraha/internal/modules/business/repository"
	categoryRepo "62teknologi-senior-backend-test-muhammad_risky_nugraha/internal/modules/category/repository"
	coordinateRepo "62teknologi-senior-backend-test-muhammad_risky_nugraha/internal/modules/coordinate/repository"
	locationRepo "62teknologi-senior-backend-test-muhammad_risky_nugraha/internal/modules/location/repository"

	"github.com/IguteChung/go-errors"
)

func GetBusinessData(ctx context.Context, queryParams *domain.GetDataQuery) ([]*domain.Business, error) {
	results, err := businessRepo.GetAllData(ctx, queryParams)

	if err != nil {
		return nil, errors.Wrap(err)
	}

	for _, business := range(results) {
		categories, err := categoryRepo.GetCategories(ctx, business.ID)
		if err != nil {
			return nil, errors.Wrap(err)
		}

		business.Category = append(business.Category, categories...)
	}
	return results, nil
}

func CreateBusinessData(ctx context.Context, payload *domain.Business) (*domain.Business, error) {
	_, err := locationRepo.CreateData(ctx, &payload.Location)

	if err != nil {
		return nil, errors.Wrap(err)
	}

	_, err = coordinateRepo.CreateData(ctx, &payload.Coordinate)

	if err != nil {
		return nil, errors.Wrap(err)
	}

	businessData, err := businessRepo.CreateData(ctx, payload)

	if err != nil {
		return nil, errors.Wrap(err)
	}

	for _, val := range payload.Category {
		val.BusinessID = businessData.ID
		_, err := categoryRepo.CreateData(ctx, val)
		if err != nil {
			return nil, errors.Wrap(err)
		}
	}

	return businessData, nil
}

func DeleteBusinessData(ctx context.Context, BusinessID int64) error {
	businessData, err := businessRepo.GetDataByID(ctx, BusinessID)
	if err != nil {
		return errors.Wrap(err)
	}

	fmt.Println(businessData.Location.ID)
	fmt.Println(businessData.Coordinate.ID)
	fmt.Println(BusinessID)

	err = locationRepo.DeleteData(ctx, businessData.Location.ID)
	if err != nil {
		return errors.Wrap(err)
	}

	err = coordinateRepo.DeleteData(ctx, businessData.Coordinate.ID)
	if err != nil {
		return errors.Wrap(err)
	}

	err = categoryRepo.DeleteData(ctx, BusinessID)
	if err != nil {
		return errors.Wrap(err)
	}

	err = businessRepo.DeleteData(ctx, BusinessID)
	if err != nil {
		return errors.Wrap(err)
	}
	return nil
}

func UpdateBusinessData(ctx context.Context, payload *domain.Business) error {
	err := locationRepo.UpdateData(ctx, &payload.Location)
	if err != nil {
		return errors.Wrap(err)
	}

	err = coordinateRepo.UpdateData(ctx, &payload.Coordinate)
	if err != nil {
		return errors.Wrap(err)
	}

	for _, val := range payload.Category {
		err = categoryRepo.UpdateData(ctx, val)
		if err != nil {
			return errors.Wrap(err)
		}
	}

	err = businessRepo.UpdateData(ctx, payload)
	if err != nil {
		return errors.Wrap(err)
	}
	return nil
}