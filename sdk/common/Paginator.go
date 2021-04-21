package common

type Paginator struct {
	PageNo      int
	HasNextPage bool
	NextID      string
	Next        func() (interface{}, error)
}

func NewPaginator(pageType string) *Paginator {
	if pageType == "number" {
		return &Paginator{
			PageNo:      1,
			NextID:      "",
			HasNextPage: true,
		}
	}
	if pageType == "cursor" {
		return &Paginator{
			PageNo:      0,
			NextID:      "*",
			HasNextPage: true,
		}
	}
	return &Paginator{}
}

func (p *Paginator) HasNext() bool {
	return p.HasNextPage
}

func (p *Paginator) SetPaginator(hasNext bool, pageNo int, nextID string) {
	p.HasNextPage = hasNext
	p.NextID = nextID
	p.PageNo = pageNo
}
