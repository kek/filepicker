package main

import (
	"reflect"
	"testing"
)

func TestListFiles(t *testing.T) {
	files := ListFiles()
	expected := []string{"go.mod", "go.sum", "main.go", "main_test.go", "tools.go"}
	if !reflect.DeepEqual(files, expected) {
		t.Errorf("got %v want %v", files, expected)
	}
}
