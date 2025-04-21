package pagination

// PageableRequestInterface defines a contract for handling pageable request data,
// including default values for sorting, pagination, and search parameters.
//
// Implementations of this interface are expected to manage default values and
// provide accessors (getters) and mutators (setters) for:
//
//   - Sort direction (asc/desc)
//   - Field to sort by
//   - Page number
//   - Items per page (limit)
//   - Search keyword
type PageableRequestInterface interface {
	SetDefaultSort(string)
	SetDefaultSortBy(string)
	SetDefaultPage(int)
	SetDefaultLimit(int)
	SetDefaultSearch(string)
	GetDefaultSort() string
	GetDefaultSortBy() string
	GetDefaultPage() int
	GetDefaultLimit() int
	GetDefaultSearch() string
}
