package lexer

import "fmt"

type Grammar struct {
	Lexemes map[string]string
}

func (g Grammar) Print() {
	for val := range g.Lexemes {
		fmt.Printf("%v:\t%v\n", val, g.Lexemes[val])
	}
}
