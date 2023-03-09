package repository

import (
	"context"

	"github.com/IguteChung/go-errors"
	"62teknologi-senior-backend-test-muhammad_risky_nugraha/internal/boot"
	"62teknologi-senior-backend-test-muhammad_risky_nugraha/internal/domain"
	"62teknologi-senior-backend-test-muhammad_risky_nugraha/internal/utils/database/builder"
)

func GetCategories(ctx context.Context, BusinessID int64) ([]*domain.Category, error) {
	db := boot.MainDBConn

	conditions := []*builder.Expression{
		builder.NewExpression("business_id", "=", BusinessID),
	}

	query, params := builder.NewBuilder("categories").
		Select("title", "alias").
		Condition(
			conditions...
		).
		Get()

	rows, err := db.Raw(query, params...).Rows()
	if err != nil {
		return nil, errors.Wrap(err)
	}

	defer rows.Close()

	results := make([]*domain.Category, 0)

	for rows.Next() {
		Categories := &domain.Category{}

		err := rows.Scan(
			&Categories.Title,
			&Categories.Alias,
		)

		if err != nil {
			return nil, errors.Wrap(err)
		}

		results = append(results, Categories)
	}

	return results, nil
}

func CreateData(ctx context.Context, payload *domain.Category) (*domain.Category, error) {
	db := boot.MainDBConn

	insertData := []*builder.SetStatement{
		{Field: "id", Value: nil},
		{Field: "title", Value: payload.Title},
		{Field: "alias", Value: payload.Alias},
		{Field: "business_id", Value: payload.BusinessID},
	}

	query, params := builder.NewBuilder("categories").
		SetData(insertData...).
		Insert()

	err := db.Exec(query, params...).Error
	if err != nil {
		return nil, errors.Wrap(err)
	}

	query, params = builder.NewBuilder("categories").
		Select("id").
		OrderByDesc("id").
		First()

	err = db.Raw(query, params...).Scan(&payload.ID).Error
		if err != nil {
			return nil, errors.Wrap(err)
		}

	return payload, nil
}

func DeleteData(ctx context.Context, BusinessID int64) error {
	db := boot.MainDBConn

	query, params := builder.NewBuilder("categories").
		Condition(
			builder.Where(builder.NewExpression("business_id", "=", BusinessID)),
		).
		Delete()

	err := db.Exec(query, params...).Error

	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}

func UpdateData(ctx context.Context, payload *domain.Category) error {
	db := boot.MainDBConn

	insertData := []*builder.SetStatement{
		{Field: "title", Value: payload.Title},
		{Field: "alias", Value: payload.Alias},
	}

	query, params := builder.NewBuilder("categories").
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