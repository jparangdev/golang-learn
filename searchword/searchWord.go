package searchword

import (
	"fmt"
	"os"
	"path/filepath"
)

func Run() {
	if len(os.Args) < 3 {
		fmt.Println("Two or more runtime arguments ar required")
		return
	}

	word := os.Args[1]
	files := os.Args[2:]
	fmt.Println("Target word:", word)
	PrintAllFiles(files)
}

func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func PrintAllFiles(files []string) {
	for _, path := range files {
		filelist, err := GetFileList(path)
		if err != nil {
			fmt.Println("it is wrong path, err:", err, " path:", path)
			continue
		}
		fmt.Println("file list")
		for _, name := range filelist {
			fmt.Println(name)
		}
	}
}
