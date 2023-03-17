package main

import (
	"fmt"
	"math"

	_ "github.com/wujiyu98/gqframe/config"
)

func getList(total uint, size uint, slot uint, currentPage uint) []string {
	var lists []string
	pageCount := uint(math.Ceil(float64(total) / float64(size)))

	if pageCount <= slot+2 {
		for i := 1; i <= int(pageCount); i++ {
			lists = append(lists, fmt.Sprint(i))
		}
		return lists

	}
	if pageCount == slot+3 {
		if currentPage < slot {
			for i := 1; i <= int(pageCount)-2; i++ {
				lists = append(lists, fmt.Sprint(i))
			}
			lists = append(lists, "...")
			lists = append(lists, fmt.Sprint(pageCount))
			return lists
		} else {
			lists = append(lists, fmt.Sprint(1))
			lists = append(lists, "...")
			space := (slot - 1) / 2
			for i := slot - space; i <= slot+space; i++ {
				lists = append(lists, fmt.Sprint(i))
			}
			lists = append(lists, fmt.Sprint(pageCount))
			return lists
		}
	}

	if currentPage < slot {
		for i := 1; i <= int(pageCount)-2; i++ {
			lists = append(lists, fmt.Sprint(i))
		}
		lists = append(lists, "...")
		lists = append(lists, fmt.Sprint(pageCount))
		return lists
	}
	if currentPage < pageCount-slot {
		lists = append(lists, fmt.Sprint(1))
		lists = append(lists, "...")
		space := (slot - 1) / 2
		for i := currentPage - space; i <= currentPage+space; i++ {
			lists = append(lists, fmt.Sprint(i))
		}
		lists = append(lists, "...")
		lists = append(lists, fmt.Sprint(pageCount))
		return lists
	}
	lists = append(lists, fmt.Sprint(1))
	lists = append(lists, "...")
	for i := int(pageCount - slot); i <= int(pageCount); i++ {
		lists = append(lists, fmt.Sprint(i))
	}

	return lists
}

func getList2(total uint, size uint, slot uint, currentPage uint) []string {
	var lists []string
	pageCount := uint(math.Ceil(float64(total) / float64(size)))
	if currentPage > pageCount {
		currentPage = pageCount

	}
	if pageCount <= slot {
		fmt.Println("case1")
		for i := 1; i <= int(pageCount); i++ {
			lists = append(lists, fmt.Sprint(i))
		}
		return lists

	}
	space := (slot - 1) / 2
	if currentPage <= space {
		fmt.Println("case2")
		for i := 1; i < int(slot-1); i++ {
			lists = append(lists, fmt.Sprint(i))
		}
		lists = append(lists, "...")
		lists = append(lists, fmt.Sprint(pageCount))
		return lists

	}
	if currentPage > pageCount-(slot-2) {
		fmt.Println("case3")
		lists = append(lists, fmt.Sprint(1))
		lists = append(lists, "...")
		for i := int(pageCount - (slot - 3)); i <= int(pageCount); i++ {
			lists = append(lists, fmt.Sprint(i))
		}
		return lists
	}
	fmt.Println("case4")
	lists = append(lists, fmt.Sprint(1))
	lists = append(lists, "...")
	for i := int(currentPage - (space - 2)); i <= int(currentPage+(space-2)); i++ {
		lists = append(lists, fmt.Sprint(i))
	}
	lists = append(lists, "...")
	lists = append(lists, fmt.Sprint(pageCount))

	return lists
}

func main() {

	fmt.Println(getList2(80, 10, 7, 5))

}
