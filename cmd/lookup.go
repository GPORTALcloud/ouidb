package main

import (
	"fmt"
	"github.com/GPORTALcloud/ouidb/pkg/ouidb"
	"os"
	"syscall"
)

//go:generate go run cmd/scripts/includedb.go
func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Printf("Missing argument\n")
		syscall.Exit(1)
	}

	db, err := ouidb.New()
	if err != nil {
		fmt.Printf("Failed to initialize db: %s\n", err.Error())
		syscall.Exit(1)
	}

	result, err := db.Lookup(args[0])
	if err != nil {
		fmt.Printf("Lookup failed: %s\n", err.Error())
		syscall.Exit(1)
	}

	fmt.Println(result)
}
