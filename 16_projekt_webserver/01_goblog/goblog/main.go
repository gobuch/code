package main

import (
	"flag"
	"fmt"
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
	Content    string
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

func makePageHandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "url: %s: Seite", r.URL.Path)
	}
}

func makeIndexHandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "url: %s: Index", r.URL.Path)
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
	p.Content = string(blackfriday.MarkdownCommon(b))
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
