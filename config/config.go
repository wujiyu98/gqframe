package config

import (
	"bytes"
	"embed"

	"github.com/spf13/viper"
)

//go:embed *.toml

var fs embed.FS

// 配置文件 直接本地目录文件名如 dev.toml
func init() {
	content, _ := fs.ReadFile("dev.toml")
	// content, _ := fs.ReadFile("pro.toml")
	viper.SetConfigType("toml")
	viper.ReadConfig(bytes.NewBuffer(content))

}
