package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// Reads all .txt files in the current folder
// and encodes them as strings literals in textfiles.go
func main() {

	resp, err := http.Get("https://gitlab.com/wireshark/wireshark/raw/master/manuf")
	if err != nil {
		log.Fatal(err.Error())
	}

	file, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	fileContent := strings.ReplaceAll(string(file), "`", "'")
	out, _ := os.Create("pkg/ouidb/database.go")
	out.Write([]byte("package ouidb \n\nconst (\n"))
	out.Write([]byte("oui = `"))
	io.Copy(out, strings.NewReader(fileContent))
	out.Write([]byte("`\n"))
	out.Write([]byte(")\n"))
}
