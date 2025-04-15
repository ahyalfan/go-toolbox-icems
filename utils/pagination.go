package utils

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

type PagebleRequest struct {
	Sort   string `query:"sort"`
	SortBy string `query:"sort_by"`
	Page   int    `query:"page"`
	Limit  int    `query:"limit"`
	Search string `query:"search"`
}

func ValidateAndPrepareRequest(pageble PagebleRequest) (PagebleRequest, error) {
	if pageble.SortBy == "" {
		pageble.SortBy = "created_at" // Default SortBy adalah created_at
	}

	if pageble.Sort == "" {
		pageble.Sort = "asc" // Default Sort adalah ascending
	} else {
		pageble.Sort = strings.ToLower(pageble.Sort)
		if pageble.Sort != "asc" && pageble.Sort != "desc" {
			return pageble, errors.New("sort must be 'asc' or 'desc'")
		}
	}

	if pageble.Page <= 0 {
		pageble.Page = 1 // Default Page adalah 1
	}

	if pageble.Limit <= 0 {
		pageble.Limit = 15 // Default Limit adalah 10
	} else if pageble.Limit > 100 {
		pageble.Limit = 100 // Batasi Limit maksimum 100
	}

	if len(pageble.Search) > 255 {
		pageble.Search = pageble.Search[:255]
	}

	return pageble, nil
}

func TotalPage(totalData int64, limit int) int {
	if totalData == 0 {
		return 0
	}
	return int(math.Ceil(float64(totalData) / float64(limit)))
}

func GenerateOffset(page, limit int) int {
	return (page - 1) * limit
}

func FormatPaginationInfo(pageble PagebleRequest) string {
	return fmt.Sprintf("Page: %d, Limit: %d, SortBy: %s, Sort: %s", pageble.Page, pageble.Limit, pageble.SortBy, pageble.Sort)
}
