package config

import (
	"bytes"
	"embed"

	"github.com/spf13/viper"
)

var Cfg *viper.Viper

//go:embed *.toml

var fs embed.FS

// 配置文件 直接本地目录文件名如 dev.toml
func Init(configFile string) {
	content, _ := fs.ReadFile(configFile)
	viper.SetConfigType("toml")
	viper.ReadConfig(bytes.NewBuffer(content))

}
