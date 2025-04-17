package pagination

type Metadata struct {
	TotalData int64 `json:"total_data"`
	TotalPage int   `json:"total_page"`
	Page      int   `json:"page"`
	Size      int   `json:"size"`
}

type PageableResponse[T any] struct {
	Data     []T      `json:"data"`
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
