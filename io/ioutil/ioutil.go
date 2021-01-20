package ioutil

import (
	"bytes"
	"io"
	"os"
)

func readAll(r io.Reader, capacity int64) (b []byte, err error) {
	var buf bytes.Buffer

	defer func() {
		// deferの中でrecover関数呼び出すと、panic関数に渡されたオブジェクトを取得できる
		e := recover()
		if e == nil {
			return
		}

		// e.(error)はinterfaceを型にキャストしている。変換できたかどうかは戻り値で判断。
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()

	// Grow関数の引数はint型なので、capacityがint型に収まるかチェックしている。
	if (int64(int(capacity))) == capacity {
		buf.Grow(int(capacity))
	}

	_, err = buf.ReadFrom(r)

	return buf.Bytes(), err
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
