package typedefs

type Tokens []string
type FileTermFreqTable map[string]int
type FileIndex struct {
	Path  *string
	Table *FileTermFreqTable
}
type DirIndex []*FileIndex
