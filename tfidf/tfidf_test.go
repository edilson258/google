package tfidf

import (
	"testing"

	"github.com/edilson258/google/typedefs"
)

func TestCalcTF(t * testing.T) {
  fileIndex := typedefs.FileIndex {
    Table: make(typedefs.FileTermFreqTable),
  }

  fileIndex.Table["hello"] = 3
  fileIndex.Table["foo"] = 34
  fileIndex.Table["bar"] = 8
  fileIndex.Table["baz"] = 2

  expextedTF := float64(8 / 4)
  foundTF := CalcTF("bar", fileIndex)

  if expextedTF != foundTF {
    t.Errorf("[ERROR]: Expected %f, found %f\n", expextedTF, foundTF)
  } 
}

func TestCalcIDF(t * testing.T) {

}
