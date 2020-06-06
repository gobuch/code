package sum

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func sum(r io.Reader, w io.Writer) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return fmt.Errorf("Error reading data: %w", err)
	}
	var numbers []int
	err = json.Unmarshal(b, &numbers)
	if err != nil {
		return fmt.Errorf("Error unmarshaling numbers: %w", err)
	}
	summe := 0
	for _, n := range numbers {
		summe = n + summe
	}
	b, err = json.Marshal(summe)
	_, err = w.Write(b)
	return err
}

func sumFile(inFileName, outFileName string) error {
	in, err := os.Open(inFileName)
	defer in.Close()
	if err != nil {
		return fmt.Errorf("sumFile: open inFile: %w", err)
	}
	out, err := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0666)
	defer out.Close()
	if err != nil {
		return fmt.Errorf("sumFile: creating outFile: %w", err)
	}
	return sum(in, out)
}
