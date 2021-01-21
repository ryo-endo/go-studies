package ioutil

import (
	"os"
	"path/filepath"
	"strings"
)

func TempDir(dir, pattern string) (name string, err error) {
	if dir == "" {
		dir = os.TempDir()
	}

	prefix, suffix := prefixAndSuffix(pattern)

	// TODO: 実装まだ
	for i := 0; i < 10000; i++ {
		try := filepath.Join(dir, prefix+"ddd"+suffix)
		err = os.Mkdir(try, 0700)
		name = try
		break
	}
	return
}

// 渡されたパターンを"*"の前後でprefixとsuffixに分割する関数
func prefixAndSuffix(pattern string) (prefix, suffix string) {
	// strings.LastIndex関数は見つかった場合は0以上を返す
	if pos := strings.LastIndex(pattern, "*"); pos != -1 {
		prefix, suffix = pattern[:pos], pattern[pos+1:]
	} else {
		prefix = pattern
	}
	return
}
