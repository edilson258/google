package fs

import (
	"os"
	"testing"
)

func TestBuildPath(t *testing.T) {
	lhs, rhs := "home", "foo"
	expectedPath := "home/foo"
	foundPath := buildPath(lhs, rhs)

	if foundPath != expectedPath {
		t.Errorf("[ERROR]: Expected %s, found %s\n", expectedPath, foundPath)
	}
}

func TestGetAllDirAndSubDirFiles(t *testing.T) {
	dir_path := "tmp_dir"

	err := os.Mkdir(dir_path, 0755)
	if err != nil {
		t.Errorf("[ERROR]: TestGetAllDirAndSubDirFiles Setup failed: %s\n", err)
	}
	defer os.RemoveAll(dir_path)

	subDirs := []string{"hello", "foo"}
	for _, subDirName := range subDirs {
		err := os.Mkdir(buildPath(dir_path, subDirName), 0755)
		if err != nil {
			t.Errorf("[ERROR]: %s\n", err)
		}
	}

	expectedFilePaths := []string{
		"hello/mom.txt",
		"hello/hello69.txt",
		"file1.txt",
		"file2.txt",
		"foo/bar.txt",
		"foo/bazz.txt",
	}

	for _, filePath := range expectedFilePaths {
		_, err := os.Create(buildPath(dir_path, filePath))
		if err != nil {
			t.Errorf("[ERROR]: %s\n", err)
		}
	}

	filePaths := GetAllDirAndSubDirFiles(dir_path)

	if len(filePaths) != len(expectedFilePaths) {
		t.Errorf("[ERROR]: Expected %d files but found %d\n", len(expectedFilePaths), len(filePaths))
	}
}
