package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// 引数walkFnは、filepath.Walk()がディレクトリを訪問するたびに呼び出される。
// return errがnilでない場合、filepath.Walk()はerrを返す。
func visit(path string, f os.FileInfo, err error) error {
	if err != nil {
		fmt.Printf("Error accessing path %q: %v\n", path, err)
		return err
	}
	// ignore .git
	if strings.HasPrefix(path, ".git") {
		return nil
	}
	fmt.Println(path)
	return nil
}

func main() {
	err := filepath.Walk(".", visit)
	if err != nil {
		fmt.Printf("Error walking the path %q: %v\n", "..", err)
	}
}
