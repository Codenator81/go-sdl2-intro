package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// get path to assets dir in current GOPATH
func AssetsPath(filename string) string {
	assetsdirectory := filepath.Dir("assets/")
	imagePath, err := filepath.Abs(assetsdirectory)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to find assets dir: %s\n", err)
		os.Exit(6)
	}
	return imagePath + "/" + filename
}
