package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
)

var (
	flagInFolder  = flag.String("in", "./", "Input-Ordner")
	flagOutFolder = flag.String("out", "", "Output-Ordner")
	flagSize      = flag.String("size", "500x500", "maximale Größe")
)

func main() {
	flag.Parse()
	size, err := parseSize(*flagSize)
	if err != nil {
		fmt.Println("Kann Größe nicht erzeugen: ", err)
		os.Exit(1)
	}
	outFolder := *flagSize
	if *flagOutFolder != "" {
		outFolder = *flagOutFolder
	}
	err = os.MkdirAll(outFolder, 0555)
	if err != nil {
		fmt.Println("Kann Zielverzeichnis nicht erzeugen: ", err)
	}
	dir, err := ioutil.ReadDir(*flagInFolder)
	if err != nil {
		fmt.Println("Fehler beim Lesen des Ordners:")
		fmt.Println(err)
		os.Exit(1)
	}
	for _, fi := range dir {
		if fi.IsDir() || !useFile(fi.Name()) {
			continue
		}
		inPath := filepath.Join(*flagInFolder, fi.Name())
		inFile, err := os.Open(inPath)
		if err != nil {
			fmt.Printf("Fehler beim Öffnen von %s: %s", inPath, err)
			continue
		}
		outPath := filepath.Join(outFolder, fi.Name())
		outFile, err := os.OpenFile(outPath, os.O_CREATE|os.O_WRONLY, 0555)
		if err != nil {
			fmt.Printf("Fehler beim Anlegen von %s: %s", outPath, err)
			inFile.Close()
			continue
		}
		err = resize(size, inFile, outFile)
		if err != nil {
			fmt.Printf("Fehler beim Verkleinern von %s: %s", inPath, err)
		}
		inFile.Close()
		outFile.Close()
	}
}

type picSize struct {
	breite int
	höhe   int
}

func parseSize(s string) (picSize, error) {
	var ps picSize
	parts := strings.Split(s, "x")
	if len(parts) != 2 {
		return ps, fmt.Errorf("%s nicht in der korrekten Form", s)
	}
	var err error
	ps.breite, err = strconv.Atoi(parts[0])
	if err != nil {
		return ps, fmt.Errorf("ps.x: %v", err)
	}
	ps.höhe, err = strconv.Atoi(parts[1])
	if err != nil {
		return ps, fmt.Errorf("ps.y: %v", err)
	}
	return ps, nil
}

func resize(ps picSize, r io.Reader, w io.Writer) error {
	img, format, err := image.Decode(r)
	if err != nil {
		return fmt.Errorf("Fehler beim Decoding: %v", err)
	}
	resized := imaging.Fit(
		img,
		ps.breite, ps.höhe,
		imaging.Lanczos,
	)
	if format != "jpeg" {
		return fmt.Errorf("Nur jpeg wird unterstützt")
	}
	return jpeg.Encode(w, resized, nil)
}

func useFile(s string) bool {
	allowed := []string{".jpg", ".jpeg"}
	ext := filepath.Ext(s)
	for _, e := range allowed {
		if strings.EqualFold(ext, e) {
			return true
		}
	}
	return false
}
