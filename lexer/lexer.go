package lexer

import (
	"github.com/edilson258/google/typedefs"
	"strings"
	"unicode"
)

type Lexer struct {
	Content string
}

func (l *Lexer) Lex() typedefs.Tokens {
	tokens := []string{}

	for l.notEOF() {
		if isalphanum(l.Content[0]) {
			cursor := 1
			for len(l.Content) > cursor && isalphanum(l.Content[cursor]) {
				cursor += 1
			}
			tokens = append(tokens, strings.ToLower(l.eatN(cursor)))
		} else {
			l.Content = l.Content[1:]
		}
	}

	return tokens
}

func (l *Lexer) notEOF() bool {
	return len(l.Content) > 0
}

func (l *Lexer) eatN(n int) string {
	token := l.Content[0:n]
	l.Content = l.Content[n:]
	return token
}

func isalphanum(char byte) bool {
	x := rune(char)
	return unicode.IsLetter(x) || unicode.IsDigit(x)
}
