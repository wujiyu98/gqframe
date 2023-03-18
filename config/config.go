package config

import (
	"bytes"
	"embed"

	"github.com/spf13/viper"
)

//go:embed *.toml

var fs embed.FS
var V *viper.Viper

// 配置文件 直接本地目录文件名如 dev.toml
func init() {
	content, _ := fs.ReadFile("dev.toml")

	// content, _ := fs.ReadFile("pro.toml")
	V = viper.New()
	V.SetConfigType("toml")
	V.ReadConfig(bytes.NewBuffer(content))

}
