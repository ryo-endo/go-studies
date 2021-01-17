package ioutil

import (
	"bytes"
	"io"
	"os"
)

func readAll(r io.Reader, capacity int64) (b []byte, err error) {
	var buf bytes.Buffer
	return buf.Bytes(), nil
}

func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// readAllに渡すキャパシティの初期値。
	// bytes.MinReadはRead関数に渡す最小のスライスサイズ。Read関数はBuffer.ReadFromから呼び出される。
	var n int64 = bytes.MinRead

	if fi, err := f.Stat(); err == nil {
		// FileInfoからサイズを取得してメモリ確保する容量を決めている。
		// FileInfo取得できなかった場合は最小容量(bytes.MinRead)がそのまま利用される。
		// TODO: FileInfo取得できないときって？
		if size := fi.Size() + bytes.MinRead; size > n {
			n = size
		}
	}

	return readAll(f, n)
}
