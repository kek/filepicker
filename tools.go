package main

import "os"

func ListFiles() []string {
	// list files in current directory
	entries, _ := os.ReadDir(".")
	var filenames []string
	for _, file := range entries {
		if file.Name()[:1] != "." && file.Name() != "filepicker" {
			filenames = append(filenames, file.Name())
		}
	}
	return filenames
}
