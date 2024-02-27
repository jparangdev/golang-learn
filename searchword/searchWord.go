package searchword

import (
	"fmt"
	"path/filepath"
)

func Run() {

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
