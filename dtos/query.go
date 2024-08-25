package dtos

type (
	QueryDTO struct {
		Page     string
		Limit    string
		Search   string
		FilterBy string
		Filter   string
		OrderBy  string
		Order    string
	}
)
