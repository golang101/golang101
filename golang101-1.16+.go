// +build go1.16

package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

//go:embed web articles/*.html
//go:embed articles/res/*.png
//go:embed articles/res/*.jpg
var allFiles embed.FS

var (
	staticFiles fs.FS
	resFiles    fs.FS
)

var staticFilesHandler, resFilesHandler = func() (http.Handler, http.Handler) {
	staticFiles, err := fs.Sub(allFiles, path.Join("web", "static"))
	if err != nil {
		panic(fmt.Sprintf("construct static file system error: %s", err))
	}
	resFiles, err := fs.Sub(allFiles, path.Join("articles", "res"))
	if err != nil {
		panic(fmt.Sprintf("construct res file system error: %s", err))
	}

	//paths1 := printFS("static files", staticFiles)
	//paths2 := printFS("res files", resFiles)
	//for _, path := range paths1 {
	//	openFileInFS(staticFiles, path)
	//}
	//for _, path := range paths2 {
	//	openFileInFS(resFiles, path)
	//}

	return http.FileServer(http.FS(staticFiles)), http.FileServer(http.FS(resFiles))
}()

func loadArticleFile(file string) ([]byte, error) {
	content, err := allFiles.ReadFile(path.Join("articles", file))
	if err != nil {
		return nil, err
	}
	return content, nil
}

func parseTemplate(commonPath string, files ...string) *template.Template {
	ts := make([]string, len(files))
	for i, f := range files {
		ts[i] = path.Join(commonPath, f)
	}
	return template.Must(template.ParseFS(allFiles, ts...))
}

func updateGolang101() {
	if _, err := os.Stat(filepath.Join(".", "golang101.go")); err == nil {
		pullGolang101Project("")
		return
	}
	if filepath.Base(os.Args[0]) == "golang101" {
		log.Println("go", "install", "go101.org/golang101@latest")
		output, err := runShellCommand(time.Minute/2, "", "go", "install", "go101.org/golang101@latest")
		if err != nil {
			log.Printf("error: %s\n%s", err, output)
		} else {
			log.Printf("done.")
		}
	}

	// no ideas how to update
}
