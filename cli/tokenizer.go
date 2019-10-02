package cli

import (
	"strings"

	"github.com/mattn/go-shellwords"
)

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

func normalizeTokens(args []string) ([]string, error) {
	ret := []string{}

	parser := shellwords.NewParser()
	// ["a b c d"] or ["a", "b", "c", "d"]
	for _, component := range args {
		parts, err := parser.Parse(component)
		if err != nil {
			return []string{}, err
		}

		for _, arg := range parts {
			v := strings.TrimSpace(arg)
			if v != "" {
				ret = append(ret, v)
			}
		}
	}

	return ret, nil
}

func (t *Tokenizer) Tokenize(args []string) error {
	tokens, err := normalizeTokens(args)
	if err != nil {
		return err
	}
	t.tokens = tokens
	t.pos = 0

	return nil
}

// TBD: This method should be renamed to 'PopToken' as destructive method.
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
