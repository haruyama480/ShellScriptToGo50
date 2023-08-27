package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type FileInfo struct {
	Path string
	Size int64
}

func main() {
	targetDir := os.Args[1]

	files, err := os.ReadDir(targetDir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if !file.IsDir() {
			stat, err := file.Info()
			if err != nil {
				panic(err)
			}
			info := FileInfo{
				Path: filepath.Join(targetDir, stat.Name()),
				Size: stat.Size(),
			}
			fmt.Printf("File: %s, Size: %d bytes\n", info.Path, info.Size)
		}
	}
}
