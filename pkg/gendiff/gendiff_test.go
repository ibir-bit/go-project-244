package gendiff

import (
	"path/filepath"
	"testing"

	"code/pkg/parser"
)

func TestGenDiff(t *testing.T) {
	basePath := filepath.Join("testdata", "fixture")
	file1 := filepath.Join(basePath, "file1.json")
	file2 := filepath.Join(basePath, "file2.json")

	data1, err := parser.ReadFile(file1)
	if err != nil {
		t.Fatal(err)
	}

	data2, err := parser.ReadFile(file2)
	if err != nil {
		t.Fatal(err)
	}

	expected := `{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}`

	result := GenDiff(data1, data2)
	if result != expected {
		t.Errorf("expected:\n%s\ngot:\n%s", expected, result)
	}
}
