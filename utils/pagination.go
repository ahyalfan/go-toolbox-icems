package utils

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/ahyalfan/go-toolbox-icems/utils/pagination"
)

func ValidateAndPrepareRequest(pageable pagination.PageableRequestInterface) (pagination.PageableRequestInterface, error) {
	if pageable.GetDefaultSortBy() == "" {
		pageable.SetDefaultSortBy("created_at") // Default SortBy adalah created_at
	}

	if pageable.GetDefaultSort() == "" {
		pageable.SetDefaultSort("asc") // Default Sort adalah ascending
	} else {
		pageable.SetDefaultSort(strings.ToLower(pageable.GetDefaultSort()))
		if pageable.GetDefaultSort() != "asc" && pageable.GetDefaultSort() != "desc" {
			return pageable, errors.New("sort must be 'asc' or 'desc'")
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

func TotalPage(totalData int64, limit int) int {
	if totalData == 0 {
		return 0
	}
	return int(math.Ceil(float64(totalData) / float64(limit)))
}

func GenerateOffset(page, limit int) int {
	return (page - 1) * limit
}

func FormatPaginationInfo(pageable pagination.PageableRequestInterface) string {
	return fmt.Sprintf("Page: %d, Limit: %d, SortBy: %s, Sort: %s", pageable.GetDefaultPage(), pageable.GetDefaultLimit(), pageable.GetDefaultSortBy(), pageable.GetDefaultSort())
}
