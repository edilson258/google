package indexer

import (
	"testing"
)

func TestBuildFileTermFreqTable(t * testing.T) {
  tokens := []string {"hello", "foo", "bar", "hello", "baz", "baz"}
  expectedTable := map[string]int {
    "hello": 2,
    "foo": 1,
    "bar": 1,
    "baz": 2,
  }
  foundTable := BuildFileTermFreqTable(tokens)

  if len(expectedTable) != len(foundTable) {
    t.Error("[ERROR]: Unexpected frequency table length")
  }

  for k, v := range expectedTable {
    freq, ok := foundTable[k]
    if !ok || freq != v {
      t.Errorf("[ERROR]: Unexpected frequency table. %s must have count: %d but found %d\n", k, v, freq)
    }
  }
}
