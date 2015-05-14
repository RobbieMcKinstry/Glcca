package lexer

import "log"
import "fmt"
import "io/ioutil"
import "encoding/json"

func Lex(toBeLexed, grammarSpec string) {
	bytes, err := ioutil.ReadFile(toBeLexed)
	if err != nil {
		log.Fatal(err)
	}
	unlexed := string(bytes)
	fmt.Println(unlexed)

	grammar := FetchGrammar(grammarSpec)
	grammar.Print()
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
