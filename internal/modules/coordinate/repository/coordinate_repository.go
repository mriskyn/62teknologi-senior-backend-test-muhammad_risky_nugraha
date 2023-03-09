package repository

import (
	"context"
	
	"github.com/IguteChung/go-errors"

	"62teknologi-senior-backend-test-muhammad_risky_nugraha/internal/boot"
	"62teknologi-senior-backend-test-muhammad_risky_nugraha/internal/domain"
	"62teknologi-senior-backend-test-muhammad_risky_nugraha/internal/utils/database/builder"

)

func CreateData(ctx context.Context, payload *domain.Coordinate) (*domain.Coordinate ,error) {
	db := boot.MainDBConn

	insertData := []*builder.SetStatement{
		{Field: "id", Value: nil},
		{Field: "latitude", Value: payload.Latitude},
		{Field: "logitude", Value: payload.Logitude},
	}

	query, params := builder.NewBuilder("coordinates").
		SetData(insertData...).
		Insert()

	err := db.Exec(query, params...).Error
	if err != nil {
		return nil, errors.Wrap(err)
	}

	query, params = builder.NewBuilder("coordinates").
		Select("id").
		OrderByDesc("id").
		First()

	err = db.Raw(query, params...).Scan(&payload.ID).Error
		if err != nil {
			return nil, errors.Wrap(err)
		}

	return payload, nil
}

func DeleteData(ctx context.Context, ID int64) error {
	db := boot.MainDBConn

	query, params := builder.NewBuilder("coordinates").
		Condition(
			builder.Where(builder.NewExpression("id", "=", ID)),
		).
		Delete()

	err := db.Exec(query, params...).Error

	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}

func UpdateData(ctx context.Context, payload *domain.Coordinate) error {
	db := boot.MainDBConn

	insertData := []*builder.SetStatement{
		{Field: "latitude", Value: payload.Latitude},
		{Field: "logitude", Value: payload.Logitude},
	}

	query, params := builder.NewBuilder("coordinates").
		SetData(insertData...).
		Condition(
			builder.Where(builder.NewExpression("id", "=", payload.ID)),
		).
		Update()

	err := db.Exec(query, params...).Error
	if err != nil {
		return errors.Wrap(err)
	}
	return nil
}