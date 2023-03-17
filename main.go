package main

import (
	"fmt"

	"github.com/wujiyu98/gqframe/tool/pagination"
)

func main() {
	p := pagination.Default()
	fmt.Println(p.GetPageUrls())
}
