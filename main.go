package main

import (
	"fmt"

	_ "github.com/wujiyu98/gqframe/config"
	"github.com/wujiyu98/gqframe/model"
	"github.com/wujiyu98/gqframe/repository"
)

func main() {
	reps := repository.New()

	var row model.Article

	reps.First(&row, 2)

	fmt.Println(row)
}
