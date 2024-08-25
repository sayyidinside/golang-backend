package helpers

import (
	"fmt"
	"golang/backend/dtos"
	"math"
	"strconv"

	"gorm.io/gorm"
)

type Pagination struct {
	Page         int         `json:"page"`
	Limit        int         `json:"limit"`
	TotalPages   int         `json:"total_pages"`
	TotalRows    int         `json:"total_rows"`
	FirstPage    string      `json:"first_page"`
	PreviousPage string      `json:"previous_page"`
	NextPage     string      `json:"next_page"`
	LastPage     string      `json:"last_page"`
	FromRow      int         `json:"from_row"`
	ToRow        int         `json:"to_row"`
	Rows         interface{} `json:"rows"`
}

func Paginate(query *dtos.QueryDTO) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(query.Page)
		if page <= 0 {
			page = 1
		}

		limit, _ := strconv.Atoi(query.Limit)
		if limit <= 0 {
			limit = 10
		}

		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}

func GeneratePaginatedQuery(query *dtos.QueryDTO, url string, totalRows int64, data []interface{}) *Pagination {
	// initilize required variable
	var nextPage, previousPage string
	var fromRow, toRow int
	totalRow := int(totalRows)

	// getting and setting page
	page, _ := strconv.Atoi(query.Page)
	if page <= 0 {
		page = 1
	}

	// getting and setting page
	limit, _ := strconv.Atoi(query.Limit)
	if limit <= 0 {
		limit = 10
	}

	// Calculate total page using totalRow [len(data)] and limit
	totalPages := int(math.Ceil(float64(totalRow) / float64(limit)))

	// Set url for first and last page
	firstPage := fmt.Sprintf("%s?page=1&limit=%d", url, limit)
	lastPage := fmt.Sprintf("%s?page=%d&limit=%d", url, totalPages, limit)

	// Set url for previous and next page
	if page > 1 {
		previousPage = fmt.Sprintf("%s?page=%d&limit=%d", url, page-1, limit)
	}
	if page < totalPages {
		nextPage = fmt.Sprintf("%s?page=%d&limit=%d", url, page+1, limit)
	}

	// Set from and to row (index)
	if page == 1 {
		fromRow = 1
		if limit > totalRow {
			toRow = totalRow
		} else {
			toRow = limit
		}
	} else if page <= totalPages {
		fromRow = ((page - 1) * limit) + 1

		if page == totalPages {
			toRow = totalRow
		} else {
			toRow = page * limit
		}
	}

	return &Pagination{
		Page:         page,
		Limit:        limit,
		TotalPages:   totalPages,
		TotalRows:    totalRow,
		FirstPage:    firstPage,
		PreviousPage: previousPage,
		NextPage:     nextPage,
		LastPage:     lastPage,
		FromRow:      fromRow,
		ToRow:        toRow,
		Rows:         data,
	}
}
