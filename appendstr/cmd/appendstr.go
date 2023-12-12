package cmd

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func AppendStrTo(args []string, appendContent string) {
	for _, arg := range args {
		if isFile(arg) {
			appendToFile(arg, appendContent)
		} else if isDir(arg) {
			appendToDir(arg, appendContent)
		}
	}
}

func appendToDir(dir string, appendContent string) {
	dirs, err := os.ReadDir(dir)
	if err != nil {
		panic(fmt.Sprintf("ReadDir(%s) failed: %s", dir, err))
	}

	for _, file := range dirs {
		if file.IsDir() {
			appendToDir(filepath.Join(dir, file.Name()), appendContent)
		} else {
			appendToFile(filepath.Join(dir, file.Name()), appendContent)
		}
	}
}

func appendToFile(file, appendContent string) {
	abs_path, err := filepath.Abs(file)
	if err != nil {
		fmt.Printf("Figure absolute path failed: %s - %s\n", file, err)
		return
	}
	if strings.HasPrefix(path.Base(abs_path), ".") {
		fmt.Printf("Ignore file: %s\n", abs_path)
		return
	}

	f, err := os.OpenFile(abs_path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Printf("Cannot open file %s: %s!\n", abs_path, err)
		return
	}

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	
	if _, err := f.Write([]byte(appendContent)); err != nil {
		panic(err)
	}

	fmt.Printf("Finished append '%s' to %s\n", appendContent, abs_path)
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

