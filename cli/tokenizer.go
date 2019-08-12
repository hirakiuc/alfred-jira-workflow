package cli

import "strings"

const (
	// Token types
	emptyToken = ""
)

type Tokenizer struct {
	tokens []string
	pos    int
}

func NewTokenizer() *Tokenizer {
	return &Tokenizer{
		tokens: []string{},
	}
}

func normalizeTokens(args []string) []string {
	ret := []string{}

	for _, component := range args {
		parts := strings.Fields(component)

		for _, arg := range parts {
			v := strings.TrimSpace(arg)
			if v != "" {
				ret = append(ret, v)
			}
		}
	}

	return ret
}

func (t *Tokenizer) Tokenize(args []string) {
	t.tokens = normalizeTokens(args)
	t.pos = 0
}

func (t *Tokenizer) NextToken() string {
	if !t.HasNextToken() {
		return emptyToken
	}

	token := t.tokens[t.pos]
	t.pos++

	return token
}

func (t *Tokenizer) RestOfTokens() []string {
	if !t.HasNextToken() {
		return []string{}
	}

	return t.tokens[t.pos:]
}

func (t *Tokenizer) HasNextToken() bool {
	return (len(t.tokens) - 1) >= t.pos
}
