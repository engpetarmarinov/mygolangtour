package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var buf bytes.Buffer
	zw := zip.NewWriter(&buf)
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling license."},
	}
	for _, file := range files {
		f, err := zw.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := zw.Close(); err != nil {
		log.Fatal(err)
	}

	filename := "std/archive/zip/test.zip"
	err := os.WriteFile(filename, buf.Bytes(), 0600)
	if err != nil {
		log.Fatal(err)
	}

	zr, err := zip.OpenReader(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer zr.Close()
	for _, file := range zr.File {
		fmt.Printf("Contents of %s:\n", file.Name)
		rc, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}

		if _, err := io.Copy(os.Stdout, rc); err != nil {
			log.Fatal(err)
		}
		rc.Close()
		fmt.Println()
	}
}
