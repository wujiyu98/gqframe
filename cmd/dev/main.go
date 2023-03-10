package main

import (
	"github.com/wujiyu98/gqframe/config"
	"github.com/wujiyu98/gqframe/router"
)

func main() {
	config.Init("dev.toml")
	router.Init().Run()

}
