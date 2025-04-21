package utils

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/ahyalfan/go-toolbox-icems/utils/pagination"
)

var (
	ErrInvalidSortDirection = errors.New("sort must be 'asc' or 'desc'")
	ErrInvalidPageNumber    = errors.New("page number must be >= 1")
	ErrInvalidPageSize      = errors.New("page size must be >= 1")
)

// ValidateAndPrepareRequest validates the given pageable request and prepares default values
// such as page number, page size, and sort direction if they are not properly set.
// It returns a potentially modified request along with an error if validation fails.
//
// This function expects the input to implement the PageableRequestInterface.
func ValidateAndPrepareRequest(pageable pagination.PageableRequestInterface) (pagination.PageableRequestInterface, error) {
	if pageable.GetDefaultSortBy() == "" {
		pageable.SetDefaultSortBy("created_at") // Default SortBy adalah created_at
	}

	if pageable.GetDefaultSort() == "" {
		pageable.SetDefaultSort("asc") // Default Sort adalah ascending
	} else {
		pageable.SetDefaultSort(strings.ToLower(pageable.GetDefaultSort()))
		if pageable.GetDefaultSort() != "asc" && pageable.GetDefaultSort() != "desc" {
			return pageable, ErrInvalidSortDirection
		}
	}

	if pageable.GetDefaultPage() <= 0 {
		pageable.SetDefaultPage(1) // Default Page adalah 1
	}

	if pageable.GetDefaultLimit() <= 0 {
		pageable.SetDefaultLimit(15) // Default Limit adalah 10
	} else if pageable.GetDefaultLimit() > 100 {
		pageable.SetDefaultLimit(100) // Batasi Limit maksimum 100
	}

	if len(pageable.GetDefaultSearch()) > 255 {
		pageable.SetDefaultSearch(pageable.GetDefaultSearch()[:255])
	}

	return pageable, nil
}

// TotalPage calculates the total number of pages based on total data and limit per page.
// If totalData is 0, it returns 0.
// It uses math.Ceil to ensure any remainder results in an additional page.
func TotalPage(totalData int64, limit int) int {
	if totalData == 0 {
		return 0
	}
	return int(math.Ceil(float64(totalData) / float64(limit)))
}

// GenerateOffset calculates the offset value for database queries based on the current page and limit.
// It returns (page - 1) * limit.
// Example: page 2 with limit 10 will return offset 10.
func GenerateOffset(page, limit int) int {
	return (page - 1) * limit
}

// FormatPaginationInfo returns a formatted string representing pagination parameters,
// including current page, limit, sort field, and sort direction.
func FormatPaginationInfo(pageable pagination.PageableRequestInterface) string {
	return fmt.Sprintf("Page: %d, Limit: %d, SortBy: %s, Sort: %s", pageable.GetDefaultPage(), pageable.GetDefaultLimit(), pageable.GetDefaultSortBy(), pageable.GetDefaultSort())
}
