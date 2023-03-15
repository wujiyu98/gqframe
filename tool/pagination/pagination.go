package pagination

import "math"

type Pagination struct {
	Total     uint
	PageCount uint
	Size      uint
	Page      uint
	Path      string
	Sort      string
	From      string
	Slot      uint
	Position  string
	Data      interface{}
}

func (p *Pagination) method() {

}

func (p *Pagination) setPageCount() {
	p.PageCount = uint(math.Ceil(float64(p.Total) / float64(p.Size)))

}
