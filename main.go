package main

import (
	"fmt"
	"github.com/edilson258/google/indexer"
)

func main() {
	root_dir_path := "DirForTesting"
	dirIndex := indexer.IndexDir(root_dir_path)

	for _, f := range dirIndex {
    fmt.Printf("[INFO]: %s Table\n", *f.Path)
		for k, v := range *f.Table {
			fmt.Println(k, "=", v)
		}
	}
}
