package dtos

type (
	QueryDTO struct {
		Page        string
		Limit       string
		Search      string
		FilterBy    string
		FilterValue string
		OrderBy     string
		Order       string
	}
)
