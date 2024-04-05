package lexer

import "testing"

func TestLex(t *testing.T) {
	content := "Hello, World!"
	l := Lexer{Content: content}
	expectedTokens := []string{"hello", "world"}
	foundTokens := l.Lex()

	if len(foundTokens) != len(expectedTokens) {
		t.Error("[ERROR]: Expected", expectedTokens, " Found", foundTokens)
	}

	for i := 0; i < len(foundTokens); i++ {
		if foundTokens[i] != expectedTokens[i] {
			t.Error("[ERROR]: Expected", expectedTokens, " Found", foundTokens)
		}
	}
}
