package token

type Token int

const (
	Add Token = iota
	Subtract
	Multiply
	Quotient
	Remainder
)

package calc

import "token"

func f(t token.Token) {
	switch t {
	case token.Add:
	case token.Subtract:
	case token.Multiply:
	default:
	}
}

func g(t token.Token) string {
	return map[token.Token]string{
		token.Add:      "add",
		token.Subtract: "subtract",
		token.Multiply: "multiply",
	}[t]
}