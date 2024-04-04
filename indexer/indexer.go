package indexer

import (
	"github.com/edilson258/google/fs"
	"github.com/edilson258/google/lexer"
	"github.com/edilson258/google/typedefs"
)

func BuildFileTermFreqTable(tokens typedefs.Tokens) typedefs.FileTermFreqTable {
	table := make(typedefs.FileTermFreqTable)

	for _, token := range tokens {
		_, ok := table[token]
		if ok {
			table[token] += 1
		} else {
			table[token] = 1
		}
	}

	return table
}

func IndexFile(path string) (result *typedefs.FileIndex) {
	fileContent := fs.ReadFileContent(path)

	if fileContent == nil {
		return nil
	}

	l := lexer.Lexer{Content: *fileContent}
	tokens := l.Lex()

	fileTermFreqTable := BuildFileTermFreqTable(tokens)

	fileIndex := typedefs.FileIndex{
		Path:  &path,
		Table: &fileTermFreqTable,
	}

	return &fileIndex
}

func IndexDir(path string) typedefs.DirIndex {
	dirIndex := []*typedefs.FileIndex{}

	for _, filePath := range fs.GetAllDirAndSubDirFiles(path) {
		fileIndex := IndexFile(filePath)

		if fileIndex != nil {
			dirIndex = append(dirIndex, fileIndex)
		}
	}

	return dirIndex
}
