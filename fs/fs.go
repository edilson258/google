package fs

import (
	"fmt"
	"os"
	"strings"
)

func GetAllDirAndSubDirFiles(root_dir_path string) []string {
	return _getAllDirAndSubDirFiles(root_dir_path, []string{})
}

func _getAllDirAndSubDirFiles(path string, entries []string) []string {
	dirEntries, err := os.ReadDir(path)

	if err != nil {
		fmt.Printf("[ERROR]: %s\n", err)
		return entries
	}

	for _, entry := range dirEntries {
		if entry.IsDir() {
			entries = append(entries, _getAllDirAndSubDirFiles(buildPath(path, entry.Name()), []string{})...)
		} else {
			entries = append(entries, buildPath(path, entry.Name()))
		}
	}

	return entries
}

func ReadFileContent(path string) (result *string) {
	if isTXTfile(path) || isMDfile(path) {
		fileBytes, err := os.ReadFile(path)

		if err != nil {
			fmt.Printf("[ERROR]: %s\n", err)
			return nil
		}

		content := string(fileBytes)
		return &content
	}

	fmt.Printf("[WARN]: Extension of %s if not supported\n", path)

	return nil
}

func buildPath(lhs string, rhs string) string {
	return lhs + "/" + rhs
}

func isTXTfile(path string) bool {
	return strings.HasSuffix(path, ".txt")
}

func isMDfile(path string) bool {
	return strings.HasSuffix(path, ".md")
}
