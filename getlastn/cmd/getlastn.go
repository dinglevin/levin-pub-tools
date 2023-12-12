package cmd

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func GetLastNFrom(args []string, num int64) {
	for _, arg := range args {
		if isFile(arg) {
			getLastNFromFile(arg, num)
		} else if isDir(arg) {
			getLastNFromDir(arg, num)
		}
	}
}

func getLastNFromDir(dir string, num int64) {
	dirs, err := os.ReadDir(dir)
	if err != nil {
		panic(fmt.Sprintf("ReadDir(%s) failed: %s", dir, err))
	}

	for _, file := range dirs {
		if file.IsDir() {
			getLastNFromDir(filepath.Join(dir, file.Name()), num)
		} else {
			getLastNFromFile(filepath.Join(dir, file.Name()), num)
		}
	}
}

func getLastNFromFile(file string, num int64) {
	abs_path, err := filepath.Abs(file)
	if err != nil {
		fmt.Printf("Figure absolute path failed: %s - %s\n", file, err)
		return
	}
	if strings.HasPrefix(path.Base(abs_path), ".") {
		fmt.Printf("Ignore file: %s\n", abs_path)
		return
	}

	f, err := os.OpenFile(abs_path, os.O_RDONLY, 0)
	if err != nil {
		fmt.Printf("Cannot open file %s: %s!\n", abs_path, err)
		return
	}

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	
	fs, err := f.Stat()
	if err != nil {
		fmt.Printf("Cannot stat file %s!\n", abs_path)
		return
	}

	offset := min(fs.Size(), num)
	f.Seek(-offset, io.SeekEnd)
	bytes := make([]byte, offset) 
	
	_, err = f.Read(bytes)
	if err != nil {
		fmt.Printf("Read last %d bytes failed from %s: %s\n", num, abs_path, err)
		return
	}

	fmt.Printf("Last %d bytes from '%s' is %s\n", num, abs_path, string(bytes[:]))
}

// 判断目录是否存在
func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func isDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func isFile(path string) bool {
	return !isDir(path)
}

func min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}