package search

import (
	"github.com/edilson258/google/lexer"
	"github.com/edilson258/google/tfidf"
	"github.com/edilson258/google/typedefs"
)

func Search(dirIndex typedefs.DirIndex, query string) typedefs.Session {
	l := lexer.Lexer{Content: query}
	session := make(typedefs.Session)

	for _, term := range l.Lex() {
		tfSession := make(typedefs.Session)

		for _, fileIndex := range dirIndex {
			tf := tfidf.CalcTF(term, *fileIndex)
			if tf > 0 {
				tfSession[*fileIndex.Path] = tf
			}
		}

		docsWithTermCount := float64(len(tfSession))
		totalDocsCount := float64(len(dirIndex))
		idf := tfidf.CalcIDF(docsWithTermCount, totalDocsCount)

		for docName, docTFScore := range tfSession {
			session[docName] = session[docName] + (docTFScore * idf)
		}
	}

	return session
}
