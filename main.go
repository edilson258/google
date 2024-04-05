package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/edilson258/google/indexer"
	"github.com/edilson258/google/search"
	"github.com/edilson258/google/typedefs"
)

func main() {
	flag.Parse()

	if len(flag.Args()) < 2 {
		fmt.Println("\n[ERROR]: Missing arguments")
		PrintUsage()
		os.Exit(1)
	}

	switch flag.Arg(0) {
	case "index":
		storeIndexDir(flag.Arg(1))
	case "search":
		if len(flag.Args()) < 3 {
			fmt.Println("\n[ERROR]: Missing arguments for search command")
			PrintSearchUsage()
		}
		searchAndDisplayResult(flag.Arg(1), flag.Arg(2))
	default:
		fmt.Printf("\n[ERROR]: Unknown command %s\n", flag.Arg(0))
		PrintUsage()
	}
}

func SortMapKeys(xs map[string]float64) []string {
	keys := make([]string, 0, len(xs))

	for key := range xs {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool { return xs[keys[i]] > xs[keys[j]] })

	return keys
}

func storeIndexDir(dirPath string) {
	dirIndex := indexer.IndexDir(dirPath)
	dirIndexJson, err := json.Marshal(dirIndex)

	if err != nil {
		fmt.Printf("[ERROR]: Couldn't convert directory index to json: %s\n", err)
		return
	}

	f, err := os.Create(dirPath + "Index.json")
	defer f.Close()

	if err != nil {
		fmt.Printf("[ERROR]: Couldn't create file to store directory index: %s\n", err)
		return
	}

	f.Write(dirIndexJson)
}

func searchAndDisplayResult(indexPath string, query string) {
	var dirIndex typedefs.DirIndex

	stream, err := os.ReadFile(indexPath)
	if err != nil {
		fmt.Printf("[ERROR]: %s", err)
		os.Exit(1)
	}

	err = json.Unmarshal(stream, &dirIndex)
	if err != nil {
		fmt.Printf("[ERROR]: Couldn't load directory index from %s: %s", indexPath, err)
		os.Exit(1)
	}

	result := search.Search(dirIndex, query)

	for _, key := range SortMapKeys(result) {
		if result[key] < 0.00001 {
			continue
		}
		fmt.Printf("%s, %f\n", key, result[key])
	}
}

func PrintUsage() {
	fmt.Println(`
Usage:

    ./google <command> [arguments]
    
The commands are:
    index  <path>             create index of a directory
    search <index>  <query>   search in directory index
  `)
}

func PrintSearchUsage() {
	fmt.Println(`
    ./google <index> <query>
    
    index         directory index
    query         term to search for
  `)
}
