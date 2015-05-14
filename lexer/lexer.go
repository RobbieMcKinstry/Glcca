package lexer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

func Lex(toBeLexed, grammarSpec string) {
	bytes, err := ioutil.ReadFile(toBeLexed)
	if err != nil {
		log.Fatal(err)
	}
	unlexed := string(bytes)
	fmt.Println(unlexed)

	grammar := FetchGrammar(grammarSpec)
	grammar.Print()

	lexemes := BuildRegexList(grammar)
	for _, lex := range lexemes {
		fmt.Printf("Name:\t%v\n", lex.Name)
	}
}

func FetchGrammar(filename string) *Grammar {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var g *Grammar = new(Grammar)
	err = json.Unmarshal(bytes, g)
	if err != nil {
		log.Fatal(err)
	}
	return g
}

type Lexeme struct {
	Name    string
	Matcher *regexp.Regexp
}

func BuildRegexList(g *Grammar) []*Lexeme {
	var (
		lexemes []*Lexeme = make([]*Lexeme, 0, len(g.Lexemes))
	)

	fmt.Printf("Number of lexemes is %v\n", len(g.Lexemes))
	for name := range g.Lexemes {
		lexeme := new(Lexeme)
		lexeme.Name = name
		lexeme.Matcher = regexp.MustCompile(g.Lexemes[name])
		lexemes = append(lexemes, lexeme)
	}
	return lexemes
}
