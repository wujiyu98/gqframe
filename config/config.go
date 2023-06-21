package config

import (
	"bytes"
	"embed"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

//go:embed *.toml

var fs embed.FS
var V *viper.Viper

// 配置文件 直接本地目录文件名如 dev.toml
func init() {
	var b []byte
	V = viper.New()
	args := os.Args

	if len(args) > 1 {
		if args[1] == "pro" {
			b, _ = fs.ReadFile("pro.toml")
			fmt.Println("use pro")
		}
	} else {
		b, _ = fs.ReadFile("dev.toml")
		fmt.Println("use dev")
	}
	V.SetConfigType("toml")
	V.ReadConfig(bytes.NewReader(b))

}
