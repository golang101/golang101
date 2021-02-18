// +build !embed

package main

import (
	"go/build"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var staticFilesHandler = http.FileServer(http.Dir(filepath.Join(rootPath, "web", "static")))
var resFilesHandler = http.FileServer(http.Dir(filepath.Join(rootPath, "articles", "res")))

func loadArticleFile(file string) ([]byte, error) {
	return ioutil.ReadFile(filepath.Join(rootPath, "articles", file))
}

func parseTemplate(commonPaths []string, files ...string) *template.Template {
	cp := filepath.Join(commonPaths...)
	ts := make([]string, len(files))
	for i, f := range files {
		ts[i] = filepath.Join(rootPath, cp, f)
	}
	return template.Must(template.ParseFiles(ts...))
}

func updateGolang101() {
	pullGolang101Project(rootPath)
}

//=================================

var rootPath = findGo101ProjectRoot()

func findGo101ProjectRoot() string {
	if _, err := os.Stat(filepath.Join(".", "golang101.go")); err == nil {
		return "."
	}

	for _, name := range []string{
		"gitlab.com/golang101/golang101", "gitlab.com/Golang101/golang101",
		"github.com/golang101/golang101", "github.com/Golang101/golang101",
	} {
		pkg, err := build.Import(name, "", build.FindOnly)
		if err == nil {
			return pkg.Dir
		}
	}

	return "."
}
