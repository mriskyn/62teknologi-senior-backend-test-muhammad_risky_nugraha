package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/IguteChung/go-errors"

	"62teknologi-senior-backend-test-muhammad_risky_nugraha/internal/boot"
	"62teknologi-senior-backend-test-muhammad_risky_nugraha/internal/domain"
	"62teknologi-senior-backend-test-muhammad_risky_nugraha/internal/utils/database/builder"
)

func CreateData(ctx context.Context, payload *domain.Location) (*domain.Location ,error) {
	db := boot.MainDBConn
	var displayAddress strings.Builder

	displayAddress.WriteString("[")
	for i, val := range payload.DisplayAddress {
		if i != 0 {
			displayAddress.WriteString(", ")
		} 

		fmt.Fprintf(&displayAddress, "\"%v\"", val)
	}
	displayAddress.WriteString("]")

	insertData := []*builder.SetStatement{
		{Field: "id", Value: nil},
		{Field: "address1", Value: payload.Address1},
		{Field: "address2", Value: payload.Address2},
		{Field: "address3", Value: payload.Address3},
		{Field: "city", Value: payload.City},
		{Field: "zip_code", Value: payload.ZipCode},
		{Field: "country", Value: payload.Country},
		{Field: "state", Value: payload.State},
		{Field: "display_address", Value: displayAddress.String()},
	}

	query, params := builder.NewBuilder("locations").
		SetData(insertData...).
		Insert()

	err := db.Exec(query, params...).Error
	if err != nil {
		return nil, errors.Wrap(err)
	}

	query, params = builder.NewBuilder("locations").
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

	query, params := builder.NewBuilder("locations").
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

func UpdateData(ctx context.Context, payload *domain.Location) error {
	db := boot.MainDBConn

	insertData := []*builder.SetStatement{
		{Field: "address1", Value: payload.Address1},
		{Field: "address2", Value: payload.Address2},
		{Field: "address3", Value: payload.Address3},
		{Field: "city", Value: payload.City},
		{Field: "country", Value: payload.Country},
		{Field: "state", Value: payload.State},
		{Field: "zip_code", Value: payload.ZipCode},
	}

	query, params := builder.NewBuilder("locations").
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