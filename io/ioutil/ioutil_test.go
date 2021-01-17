package ioutil

import (
	"testing"
)

func TestReadFile(t *testing.T) {
	filename := "hogehoge"
	_, err := ReadFile(filename)
	if err == nil {
		t.Fatalf("Readfile #{filename}: error expected, none found")
	}

	filename = "ioutil_test.go"
	_, err = ReadFile(filename)
	if err != nil {
		t.Fatalf("ReadFile #{filename}: #{err}")
	}
}
