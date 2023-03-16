package main

import (
	_ "github.com/wujiyu98/gqframe/config"
	"github.com/wujiyu98/gqframe/router"
)

func main() {

	router.Init().Run()

}
