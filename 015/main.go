package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	targetDir := os.Args[1]

	cnt := make(map[string]int)
	filepath.Walk(targetDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			ext := filepath.Ext(info.Name())
			if ext == "" {
				ext = "other"
			}
			cnt[ext]++
		}
		return nil
	})

	// sort by counts
	keys := make([]string, 0, len(cnt))
	for k := range cnt {
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return cnt[keys[i]] > cnt[keys[j]]
	})

	for _, k := range keys {
		fmt.Printf("Extension: %s, Count: %d\n", k, cnt[k])
	}
}
