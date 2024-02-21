package main

import (
	"hash/adler32"
	"unicode"
)

type Token struct {
	value  string
	kind   TokenKind
	length int
}

type TokenKind int

const (
	Unknown TokenKind = iota
	Number
	Operator
	Identifier
	String
	Whitespace
	EOF
	Error
	Comment
	Keyword
	Symbol
	Punctuator
	UnknownPunctuator
	KeywordPunctuator
	SymbolPunctuator
	WhitespacePunctuator
)

func tokenize(input string) []Token {
	tokens := lex([]rune(input))
	tokens = append(tokens, Token{value: "EOF", kind: EOF, length: 0})
	return tokens
}

func eatwhitespace(input []rune, whitespace string) []rune {
	for len(input) > 0 && unicode.IsSpace(input[0]) {
		input = input[1:]
	}
	return input
}

func input() {

}

func lan(int, []Token) {
	if len(input) == 0 {
		return
	}
}

func lexToken(input []rune) *Token {
	if len(input) == 0 {
		return nil
	}
	if unicode.IsSpace(input[0]) {
		return &Token{
			value:  " ",
			length: 1,
		}
	}
	return nil
}

func lex(input []rune) (tokens []Token) {
	curser := 0
	for len(input) > 0 {
		token := lexToken(input)
		if token != nil {
			tokens = append(tokens, *token)
			input = input[token.length:]
			curser += token.length
			continue
		}
		if len(input) == 0 {
			break
		}
		tokens = append(tokens, Token{
			value:  string(input[0]),
			length: 1,
		})
		input = input[token.length:]
		curser += token.length
	}
	return tokens
}
