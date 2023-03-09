package server

type DeleteBusinessPayload struct {
	BusinessID	int64		`json:"business_id" validate:"required"`
}

type GetDataQuery struct {
	OpenNow 	bool 	`query:"open_now"`
	Limit			int		`query:"limit"`
	Offset 		int		`query:"offset"`
	Location	string`query:"location"`
	SortBy		string`query:"sort_by"`
}