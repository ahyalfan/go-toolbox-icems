package pagination

// Metadata contains the pagination details for a response,
// including the total number of items (TotalData), total pages (TotalPage),
// current page (Page), and the size (Size) of each page.
type Metadata struct {
	// TotalData is the total number of items across all pages.
	TotalData int64 `json:"total_data"`

	// TotalPage is the total number of pages available, based on the limit per page.
	TotalPage int `json:"total_page"`

	// Page is the current page being requested.
	Page int `json:"page"`

	// Size is the number of items per page.
	Size int `json:"size"`
}

// PageableResponse is a generic structure used to represent a paginated response
// with data of type T and associated metadata such as total data, total pages,
// current page, and page size.
type PageableResponse[T any] struct {
	// Data is the paginated data for the current page.
	// It holds a slice of type T, which represents the actual items being requested.
	Data []T `json:"data"`

	// Metadata contains pagination information for the response, including total data, total pages,
	// the current page, and the number of items per page.
	Metadata Metadata `json:"metadata"`
}

type PageableRequest struct {
	Sort   string `query:"sort"`
	SortBy string `query:"sort_by"`
	Page   int    `query:"page"`
	Limit  int    `query:"limit"`
	Search string `query:"search"`
}

func (p *PageableRequest) SetDefaultSort(v string) {
	p.Sort = v
}

func (p *PageableRequest) SetDefaultSortBy(v string) {
	p.SortBy = v
}

func (p *PageableRequest) SetDefaultPage(v int) {
	p.Page = v
}

func (p *PageableRequest) SetDefaultLimit(v int) {
	p.Limit = v
}

func (p *PageableRequest) SetDefaultSearch(v string) {
	p.Search = v
}

func (p *PageableRequest) GetDefaultSort() string {
	return p.Sort
}

func (p *PageableRequest) GetDefaultSortBy() string {
	return p.SortBy
}

func (p *PageableRequest) GetDefaultPage() int {
	return p.Page
}

func (p *PageableRequest) GetDefaultLimit() int {
	return p.Limit
}

func (p *PageableRequest) GetDefaultSearch() string {
	return p.Search
}
