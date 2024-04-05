package main

import (
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/edilson258/google/indexer"
	"github.com/edilson258/google/search"
)

func PrintUsage() {
  fmt.Println("Usage:")
  fmt.Println("\t./google [dir path] [query]")
  fmt.Println("Example:")
  fmt.Println("\t./google TestDir \"what is java\" ")
}

func SortMapKeys(xs map[string]float64) []string {
	keys := make([]string, 0, len(xs))

	for key := range xs {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool { return xs[keys[i]] > xs[keys[j]] })

  return keys
}

func main() {
	flag.Parse()

	if len(flag.Args()) != 2 {
		fmt.Println("\n[ERROR]: Incorrect arguments")
		PrintUsage()
    os.Exit(1)
	}

	dirPath := flag.Arg(0)
	query := flag.Arg(1)

	dirIndex := indexer.IndexDir(dirPath)
	result := search.Search(dirIndex, query)

  
	for _, key := range SortMapKeys(result) {
		fmt.Printf("%s, %f\n", key, result[key])
	}
}
