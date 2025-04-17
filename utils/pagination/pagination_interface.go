package pagination

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
