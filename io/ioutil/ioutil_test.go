package ioutil

import (
	"os"
	"testing"
)

func checkSize(t *testing.T, path string, size int64) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		// %q 文字列表記(ダブルクオーテーション付き)
		t.Fatalf("Stat %q (looking for size %d: %s", path, size, err)
	}
	if fileInfo.Size() != size {
		t.Errorf("Stat %q: size %d want %d", path, fileInfo.Size(), size)
	}
}

func TestReadFile(t *testing.T) {
	filename := "hogehoge"
	_, err := ReadFile(filename)
	if err == nil {
		t.Fatalf("Readfile %s: error expected, none found", filename)
	}

	filename = "ioutil_test.go"
	contents, err := ReadFile(filename)
	if err != nil {
		t.Fatalf("ReadFile %s: %v", filename, err)
	}

	checkSize(t, filename, int64(len(contents)))
}
