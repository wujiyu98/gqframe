package pagination

import (
	"math"
)

type Pagination struct {
	Total uint
	Size  uint
	Page  uint
	Path  string
	Slot  uint
}

func Defalut() *Pagination {

	return &Pagination{
		Total: 100,
		Size:  10,
		Path:  "/",
	}
}

func (p *Pagination) method() {

}

func (p *Pagination) getPageCount() uint {
	if p.Total == 0 {
		return 1
	}
	return uint(math.Ceil(float64(p.Total) / float64(p.Size)))
}

func (p *Pagination) getList() {
	pageCount := p.getPageCount()
	sort := p.Slot
	if pageCount <= sort+2 {

	}

}
