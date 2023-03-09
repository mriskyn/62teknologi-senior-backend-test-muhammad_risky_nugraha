package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"62teknologi-senior-backend-test-muhammad_risky_nugraha/internal/boot"
	"62teknologi-senior-backend-test-muhammad_risky_nugraha/internal/domain"
	"62teknologi-senior-backend-test-muhammad_risky_nugraha/internal/helpers/slicehelper"
	"62teknologi-senior-backend-test-muhammad_risky_nugraha/internal/utils/database/builder"

	"github.com/IguteChung/go-errors"
)

func GetAllData(ctx context.Context, queryParams *domain.GetDataQuery) ([]*domain.Business, error) {
	db := boot.MainDBConn

	joinStmt := []*builder.JoinStatement{
		{JoinTable: "coordinates co", Condition: "b.coordinate_id=co.id", JoinType: builder.LeftJoin},
		{JoinTable: "locations lo", Condition: "b.location_id=lo.id", JoinType: builder.LeftJoin},
	}

	conditions := []*builder.Expression{}

	if queryParams.Location != "" {
		conditions = append(
			conditions, 
			builder.Where(builder.NewExpression("lo.country", "=", queryParams.Location)),
		)
	}

	if queryParams.OpenNow != nil {
		isOpen := *queryParams.OpenNow

		conditions = append(
			conditions, 
			builder.Where(builder.NewExpression("is_closed", "=", !isOpen)),
		)
	}

	orderStatement := ""
	if queryParams.SortBy != "" {
		switch queryParams.SortBy {
			case "rating":
				orderStatement = "ORDER BY rating DESC"
			case "review_count":
				orderStatement = "ORDER BY review_count DESC"
			case "distance":
				orderStatement = "ORDER BY distance ASC"
		}
	}


	var priceConditional []string
	if len(queryParams.Price) > 0 {
		for _, val := range queryParams.Price {
			priceString := strings.Repeat("$", val)
			priceConditional = append(priceConditional, priceString)
		}
		conditions = append(
			conditions, 
			builder.WhereIn("b.price", slicehelper.SliceStringToInterface(priceConditional)),
		)
	} 

	query, params := builder.NewBuilder("business b").
		Select(
			"b.id",
			"b.alias",
			"b.name",
			"b.image_url",
			"b.is_closed",
			"b.url",
			"b.review_count",
			"b.rating",
			"co.latitude",
			"co.logitude",
			"b.transactions",
			"lo.address1",
			"lo.address2",
			"lo.address3",
			"lo.city",
			"lo.country",
			"lo.state",
			"lo.zip_code",
			"lo.display_address",
			"b.price",
			"b.phone",
			"b.display_phone",
			"b.distance",
			).
		Join(joinStmt...).
		Condition(conditions...).
		Limit(queryParams.Limit).
		Offset(queryParams.Offset).
		RawOrderBy(orderStatement).
		Get()

		fmt.Println(query)
		fmt.Println()

	rows, err := db.Raw(query, params...).Rows()
	if err != nil {
		return nil, errors.Wrap(err)
	}

	defer rows.Close()

	results := make([]*domain.Business, 0)

	for rows.Next() {
		Businesses := &domain.Business{}
		err := rows.Scan(
			&Businesses.ID,
			&Businesses.Alias,
			&Businesses.Name,
			&Businesses.ImageUrl,
			&Businesses.IsClosed,
			&Businesses.Url,
			&Businesses.ReviewCount,
			&Businesses.Rating,
			&Businesses.Coordinate.Latitude,
			&Businesses.Coordinate.Logitude,
			&Businesses.TransactionsDB,
			&Businesses.Location.Address1,
			&Businesses.Location.Address2,
			&Businesses.Location.Address3,
			&Businesses.Location.City,
			&Businesses.Location.Country,
			&Businesses.Location.State,
			&Businesses.Location.ZipCode,
			&Businesses.Location.DisplayAddressDB,
			&Businesses.Price,
			&Businesses.Phone,
			&Businesses.DisplayPhone,
			&Businesses.Distance,
		)

		var unmarshalTransactions []string
		var unmarshalDisplayAddress [] string
		_ = json.Unmarshal(Businesses.TransactionsDB, &unmarshalTransactions)
		_ = json.Unmarshal(Businesses.Location.DisplayAddressDB, &unmarshalDisplayAddress)

		Businesses.Transactions = unmarshalTransactions
		Businesses.Location.DisplayAddress = unmarshalDisplayAddress

		if err != nil {
			return nil, errors.Wrap(err)
		}

		results = append(results, Businesses)
	}

	return results, nil
}

func GetDataByID(ctx context.Context, ID int64) (*domain.Business, error) {
	db := boot.MainDBConn

	result := domain.Business{}

	joinStmt := []*builder.JoinStatement{
		{JoinTable: "coordinates co", Condition: "b.coordinate_id=co.id", JoinType: builder.LeftJoin},
		{JoinTable: "locations lo", Condition: "b.location_id=lo.id", JoinType: builder.LeftJoin},
	}

	query, params := builder.NewBuilder("business b").
		Select(
			"b.id",
			"b.alias",
			"b.name",
			"b.image_url",
			"b.is_closed",
			"b.url",
			"b.review_count",
			"b.rating",
			"co.id",
			"co.latitude",
			"co.logitude",
			"b.transactions",
			"lo.id",
			"lo.address1",
			"lo.address2",
			"lo.address3",
			"lo.city",
			"lo.country",
			"lo.state",
			"lo.zip_code",
			"lo.display_address",
			"b.price",
			"b.phone",
			"b.display_phone",
			"b.distance",
			).
		Join(joinStmt...).
		Get()

	err := db.Raw(query, params...).Row().Scan(
		&result.ID,
			&result.Alias,
			&result.Name,
			&result.ImageUrl,
			&result.IsClosed,
			&result.Url,
			&result.ReviewCount,
			&result.Rating,
			&result.Coordinate.ID,
			&result.Coordinate.Latitude,
			&result.Coordinate.Logitude,
			&result.TransactionsDB,
			&result.Location.ID,
			&result.Location.Address1,
			&result.Location.Address2,
			&result.Location.Address3,
			&result.Location.City,
			&result.Location.Country,
			&result.Location.State,
			&result.Location.ZipCode,
			&result.Location.DisplayAddressDB,
			&result.Price,
			&result.Phone,
			&result.DisplayPhone,
			&result.Distance,
	)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	var unmarshalTransactions []string
	var unmarshalDisplayAddress [] string
	_ = json.Unmarshal(result.TransactionsDB, &unmarshalTransactions)
	_ = json.Unmarshal(result.Location.DisplayAddressDB, &unmarshalDisplayAddress)
	
	result.Transactions = unmarshalTransactions
	result.Location.DisplayAddress = unmarshalDisplayAddress

	return &result, nil
}

func CreateData(ctx context.Context, payload *domain.Business) (*domain.Business ,error) {
	db := boot.MainDBConn

	insertData := []*builder.SetStatement{
		{Field: "id", Value: nil},
		{Field: "alias", Value: payload.Alias},
		{Field: "name", Value: payload.Name},
		{Field: "image_url", Value: payload.ImageUrl},
		{Field: "is_closed", Value: payload.IsClosed},
		{Field: "url", Value: payload.Url},
		{Field: "review_count", Value: payload.ReviewCount},
		{Field: "rating", Value: payload.Rating},
		{Field: "coordinate_id", Value: payload.Coordinate.ID},
		{Field: "transactions", Value: payload.TransactionsDB},
		{Field: "location_id", Value: payload.Location.ID},
		{Field: "price", Value: payload.Price},
		{Field: "phone", Value: payload.Phone},
		{Field: "display_phone", Value: payload.DisplayPhone},
		{Field: "distance", Value: payload.Distance},
	}

	query, params := builder.NewBuilder("business").
		SetData(insertData...).
		Insert()
	
	
	err := db.Exec(query, params...).Error
	if err != nil {
		return nil, errors.Wrap(err)
	}

	query, params = builder.NewBuilder("business").
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

	query, params := builder.NewBuilder("business").
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

func UpdateData(ctx context.Context, payload *domain.Business) error {
	db := boot.MainDBConn

	insertData := []*builder.SetStatement{
		{Field: "alias", Value: payload.Alias},
		{Field: "name", Value: payload.Name},
		{Field: "image_url", Value: payload.ImageUrl},
		{Field: "is_closed", Value: payload.IsClosed},
		{Field: "url", Value: payload.Url},
		{Field: "review_count", Value: payload.ReviewCount},
		{Field: "rating", Value: payload.Rating},
		{Field: "transactions", Value: payload.TransactionsDB},
		{Field: "price", Value: payload.Price},
		{Field: "phone", Value: payload.Phone},
		{Field: "display_phone", Value: payload.DisplayPhone},
		{Field: "distance", Value: payload.Distance},
	}

	query, params := builder.NewBuilder("business").
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