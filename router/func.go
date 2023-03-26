package router

import (
	"fmt"
	"html/template"

	"github.com/wujiyu98/gqframe/config"
)

var funcMap = template.FuncMap{
	"asset": asset,
}

// 取assetUrl,exmaple http://localhost/
func asset(path string) string {
	host := config.V.GetString("assetUrl")
	return fmt.Sprint(host, path)
}
