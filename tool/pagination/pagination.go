package pagination

import (
	"fmt"
	"math"
	"strings"
)

type Paginator struct {
	Total       int
	PerPage     int
	CurrentPage int
	Path        string
	EachSide    int
}

type PageUrl struct {
	Page string
	Url  string
}

func (p *Paginator) method() {

}

func (p *Paginator) getCount() (count int) {
	if p.Total > 0 {
		count = int(math.Ceil(float64(p.Total) / float64(p.getPerPage())))
	}

	return
}

func (p *Paginator) GetLists() (lists []string) {
	count := p.getCount()
	eachSide := p.EachSide
	currentPage := p.CurrentPage
	if currentPage > count {
		currentPage = count
	}
	if count <= 2*eachSide+1 {
		fmt.Println("case1")
		for i := 1; i <= count; i++ {
			lists = append(lists, fmt.Sprint(i))
		}
		return
	}
	if currentPage <= eachSide {
		fmt.Println("case2")
		for i := 1; i < 2*eachSide; i++ {
			lists = append(lists, fmt.Sprint(i))
		}
		lists = append(lists, "...")
		lists = append(lists, fmt.Sprint(count))
		return
	}
	if currentPage > count-eachSide-1 {
		fmt.Println("case3")
		lists = append(lists, fmt.Sprint(1))
		lists = append(lists, "...")
		for i := count - eachSide - 1; i <= count; i++ {
			lists = append(lists, fmt.Sprint(i))
		}
		return
	}
	fmt.Println("case4")
	lists = append(lists, fmt.Sprint(1))
	lists = append(lists, "...")
	for i := currentPage - (eachSide - 2); i <= currentPage+(eachSide-2); i++ {
		lists = append(lists, fmt.Sprint(i))
	}
	lists = append(lists, "...")
	lists = append(lists, fmt.Sprint(count))

	return
}

func (p *Paginator) getPerPage() int {
	if p.PerPage > 100 {
		return 100
	}
	return p.PerPage

}

func (p *Paginator) isFirstPage(page string) bool {
	return page == "1"
}

func (p *Paginator) getUrl(page string) string {
	if p.isFirstPage(page) {
		return p.Path
	}
	if strings.Contains(p.Path, "?") {
		return p.Path + "&page=" + page
	} else {
		return p.Path + "?page=" + page
	}

}

func (p *Paginator) GetPageUrls() (pageUrls []PageUrl) {
	lists := p.GetLists()
	for _, list := range lists {
		var pageUrl PageUrl
		pageUrl.Page = list
		if list != "..." {
			p.getUrl(list)
		}
		pageUrls = append(pageUrls, pageUrl)
	}
	return
}
