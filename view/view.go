package view

import (
	"embed"
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
)

//go:embed layouts/*.html components/*html pages/*.html
var fs embed.FS

func Test() {
	layout := template.Must(template.ParseFS(fs, "layouts/master.html", "layouts/header.html", "layouts/footer.html"))
	t1 := template.Must(layout.Clone())
	t1.ParseFS(fs, "pages/message.html", "includes/message.html")
	for _, v := range t1.Templates() {
		fmt.Println(v.Name())

	}
	t1.Execute(os.Stdout, nil)

}

// fh直接引用目录文件模式
func LoadTemplates() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts := []string{"view/layouts/master.html", "view/layouts/header.html", "view/layouts/footer.html"}

	components, err := filepath.Glob("view/components/*.html")
	if err != nil {
		panic(err.Error())
	}
	pages, err := filepath.Glob("view/pages/*.html")
	if err != nil {
		panic(err.Error())
	}
	// Generate our templates map from our layouts/ and includes/ directories
	for _, page := range pages {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		layoutCopy = append(layoutCopy, page)
		files := append(layoutCopy, components...)
		r.AddFromFiles(filepath.Base(page), files...)
	}
	return r
}

func getComponents() (lists []string) {
	rows, _ := fs.ReadDir("components")
	for _, row := range rows {
		pathname := fmt.Sprintf("components/%s", row.Name())
		lists = append(lists, pathname)
	}
	return

}

// 缓存文件模式
func LoadTemplatesFs() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	layouts := []string{"layouts/master.html", "layouts/header.html", "layouts/footer.html"}
	components := getComponents()
	rows, _ := fs.ReadDir("pages")
	for _, row := range rows {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		pathname := fmt.Sprintf("pages/%s", row.Name())
		layoutCopy = append(layoutCopy, pathname)
		layoutCopy = append(layoutCopy, components...)
		t := template.Must(template.ParseFS(fs, layoutCopy...))
		r.Add(row.Name(), t)
	}
	return r
}
