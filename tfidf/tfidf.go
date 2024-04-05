package tfidf

import (
	"math"

	"github.com/edilson258/google/typedefs"
)

func CalcTF(term string, doc typedefs.FileIndex) float64 {
	freq, ok := doc.Table[term]

	if !ok {
		return 0
	}

	return float64(freq) / float64(len(doc.Table))
}

func CalcIDF(docsWithTermCount float64, docsCount float64) float64 {
	return math.Log10(docsCount / (docsWithTermCount))
}
