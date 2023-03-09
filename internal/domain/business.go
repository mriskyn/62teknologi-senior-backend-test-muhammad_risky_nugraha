package domain

type Business struct {
	ID						int64 			`json:"id" db:"id"`
	Alias					string			`json:"alias" db:"alias"`
	Name					string			`json:"name" db:"name"`
	ImageUrl			string			`json:"image_url" db:"image_url"`
	IsClosed			bool				`json:"is_closed" db:"is_closed"`
	Url						string			`json:"url" db:"url"`
	ReviewCount		int					`json:"review_count" db:"review_count"`
	Category 			[]*Category `json:"category,omitempty"`
	Rating				float64 		`json:"rating" db:"rating"`
	Coordinate 		Coordinate 	`json:"coordinate"`
	TransactionsDB[]byte			`json:"-" db:"transactions"`
	Transactions 	[]string 		`json:"transactions"`
	Location 			Location 		`json:"location"`
	Price					string			`json:"price" db:"price"`
	Phone					string			`json:"phone" db:"phone"`
	DisplayPhone	string			`json:"display_phone" db:"display_phone"`
	Distance 			float64			`json:"distance" db:"distance"`
}

type Category struct {
	ID					int64			`json:"id,omitempty" db:"id"`
	Alias				string		`json:"alias" db:"alias"`
	Title				string		`json:"title" db:"title"`
	BusinessID 	int64			`json:"-" db:"business_id"`
}

type Coordinate struct {
	ID 				int64		`json:"id,omitempty" db:"id"`
	Latitude 	float64	`json:"latitude" db:"latitude"`
	Logitude 	float64 `json:"longitude" db:"logitude"`
}

type Location struct {
	ID 							int64 	`json:"id,omitempty" db:"id"`
	Address1 				string	`json:"address1" db:"address1"`
	Address2 				string 	`json:"address2" db:"address2"`
	Address3 				string 	`json:"address3" db:"address3"`
	City 						string	`json:"city" db:"city"`
	ZipCode					string	`json:"zip_code" db:"zip_code"`
	Country 				string	`json:"country"  db:"country"`
	State 					string 	`json:"state" db:"state"`
	DisplayAddressDB[]byte  `json:"-" db:"display_address"`
	DisplayAddress	[]string`json:"display_address"`
}

type GetDataQuery struct {
	OpenNow 	*bool `query:"open_now"`
	Limit			int		`query:"limit"`
	Offset 		int		`query:"offset"`
	Location	string`query:"location"`
	SortBy		string`query:"sort_by"`
	Price			[]int	`query:"price"`
}

type DeleteBusinessPayload struct {
	BusinessID	int64		`json:"business_id" validate:"required"`
}