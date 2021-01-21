package ioutil

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestTempFile(t *testing.T) {
	dir, err := TempDir("", "TestTempFile_BadDir")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	nonexistentDir := filepath.Join(dir, "_not_exists_")
	// TODO: TempFile関数を独自で書き直す
	f, err := ioutil.TempFile(nonexistentDir, "foo")
	if f != nil || err == nil {
		t.Errorf("TempFile(%q, `foo`) = %v, %v", nonexistentDir, f, err)
	}
}
