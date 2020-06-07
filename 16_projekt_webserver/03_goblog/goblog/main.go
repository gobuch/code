package main

import (
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/russross/blackfriday"
)

type Page struct {
	Title      string
	LastChange time.Time
	Content    template.HTML
}

type Pages []Page

var (
	flagSrcFolder  = flag.String("src", "./seiten/", "blog folder")
	flagTmplFolder = flag.String("tmpl", "./templates/", "template folder")
)

func main() {
	http.HandleFunc("/page/", makePageHandlerFunc())
	http.HandleFunc("/", makeIndexHandlerFunc())
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}

func loadPage(fpath string) (Page, error) {
	var p Page
	fi, err := os.Stat(fpath)
	if err != nil {
		return p, fmt.Errorf("loadPage: %w", err)
	}
	p.Title = fi.Name()
	p.LastChange = fi.ModTime()
	b, err := ioutil.ReadFile(fpath)
	if err != nil {
		return p, fmt.Errorf("loadPage.ReadFile: %w", err)
	}
	p.Content = template.HTML(blackfriday.MarkdownCommon(b))
	return p, nil
}

func loadPages(src string) (Pages, error) {
	var ps Pages
	fs, err := ioutil.ReadDir(src)
	if err != nil {
		return ps, fmt.Errorf("loadPages.ReadDir: %w", err)
	}
	for _, f := range fs {
		if f.IsDir() {
			continue
		}
		fpath := filepath.Join(src, f.Name())
		p, err := loadPage(fpath)
		if err != nil {
			return ps, fmt.Errorf("loadPages.loadPage: %w", err)
		}
		ps = append(ps, p)
	}
	return ps, nil
}

func renderPage(w io.Writer, data interface{}, content string) error {
	tmpl, err := template.ParseFiles(
		filepath.Join(*flagTmplFolder, "base.tmpl.html"),
		filepath.Join(*flagTmplFolder, "header.tmpl.html"),
		filepath.Join(*flagTmplFolder, "footer.tmpl.html"),
		filepath.Join(*flagTmplFolder, content),
	)
	if err != nil {
		return fmt.Errorf("renderPage.ParseFiles: %w", err)
	}
	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		return fmt.Errorf("renderPage.ExecuteTemplate: %w", err)
	}
	return nil
}

func makeIndexHandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ps, err := loadPages(*flagSrcFolder)
		if err != nil {
			fmt.Println(err)
		}
		err = renderPage(w, ps, "index.tmpl.html")
		if err != nil {
			fmt.Println(err)
		}
	}
}

func makePageHandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f := r.URL.Path[len("/page/"):]
		fpath := filepath.Join(*flagSrcFolder, f)
		p, err := loadPage(fpath)
		if err != nil {
			fmt.Println(err)
		}
		err = renderPage(w, p, "page.tmpl.html")
		if err != nil {
			fmt.Println(err)
		}
	}
}
