package helper

import (
	"database/sql"
	"fmt"
	"math"
)

type PaginationCollectionPost struct {
	NextPage     int
	PreviousPage int
	CurrentPage  int
	TotalPages   int
	TwoAfter     int
	TwoBelow     int
	Offset 		int
}

func GetPaginationData(db *sql.DB, tableName string, page int, perPage int ) PaginationCollectionPost {


	var totalRows int

	queryCount := fmt.Sprintf("SELECT COUNT(*) FROM %s;", tableName)
	err := db.QueryRow(queryCount).Scan(&totalRows)
	if err != nil {
		return PaginationCollectionPost{}
	}
	totalPages := math.Ceil(float64(totalRows) / float64(perPage))
	offset := (page - 1) * 16


	return PaginationCollectionPost{
		NextPage:     page + 1,
		PreviousPage: page - 1,
		CurrentPage:  page,
		TotalPages:   int(totalPages),
		TwoAfter:     page + 2,
		TwoBelow:     page - 2,
		Offset: 	  offset,
	}
}