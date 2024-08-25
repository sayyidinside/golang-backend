package helpers

import (
	"fmt"
	"golang/backend/dtos"

	"gorm.io/gorm"
)

func Order(query *dtos.QueryDTO, allowedFields map[string]string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// Ordering logic
		orderBy := query.OrderBy
		order := query.Order

		// Validate the order_by field and retrieve the corresponding database field
		dbField, isValidOrderField := allowedFields[orderBy]

		if isValidOrderField {
			if order != "asc" && order != "desc" {
				order = "asc"
			}
			db = db.Order(fmt.Sprintf("%s %s", dbField, order))
		}

		return db
	}
}
