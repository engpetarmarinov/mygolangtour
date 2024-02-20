package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling license."},
	}
	for _, file := range files {
		header := &tar.Header{
			Name: file.Name,
			Size: int64(len(file.Body)),
			Mode: 0600,
		}

		if err := tw.WriteHeader(header); err != nil {
			log.Fatal(err)
		}

		if _, err := tw.Write([]byte(file.Body)); err != nil {
			log.Fatal(err)
		}
	}

	if err := tw.Close(); err != nil {
		log.Fatal(err)
	}

	filename := "std/archive/tar/test.tar"
	err := os.WriteFile(filename, buf.Bytes(), 0600)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	tr := tar.NewReader(bytes.NewBuffer(file))
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Contents of %s:\n", header.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatal(err)
		}
		fmt.Println()
	}
}
