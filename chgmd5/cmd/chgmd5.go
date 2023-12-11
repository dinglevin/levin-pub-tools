package cmd

import (
	"fmt"
	"os"
)

func Chgmd5(args []string, appendStr string) {

}

func appendToFiles(files []string, appendStr string) {
	for _, file := range files {
		appendToFile(file, appendStr)
	}
}

func appendToFile(file, appendStr string) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Printf("Cannot open file %s!\n", file)
		return
	}

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()
	
	if _, err := f.Write([]byte(appendStr)); err != nil {
		panic(err)
	}
}