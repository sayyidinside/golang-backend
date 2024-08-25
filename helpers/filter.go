package helpers

import (
	"fmt"
	"golang/backend/dtos"

	"gorm.io/gorm"
)

func Filter(query *dtos.QueryDTO, allowedFields map[string]string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		filterBy := query.FilterBy
		filterValue := query.FilterValue

		// Validate the filter_by field and retrieve the corresponding database field
		dbField, isValidFilterField := allowedFields[filterBy]

		if isValidFilterField && filterValue != "" {
			// Apply the filter to the query
			db = db.Where(fmt.Sprintf("%s = ?", dbField), filterValue)
		}

		return db
	}
}
