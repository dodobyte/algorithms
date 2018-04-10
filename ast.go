package main

import (
	"bufio"
	"fmt"
	"os"
)

type symbol struct {
	c   byte
	val int
}

type ast struct {
	symbol
	args []*ast
}

var sym symbol
var reader *bufio.Reader

// Lexical Analyzer

func next() {
	c, err := reader.ReadByte()
	if err != nil {
		return
	}
	if c == ' ' {
		next()
		return
	}
	if c >= '0' && c <= '9' {
		reader.UnreadByte()
		fmt.Fscanf(reader, "%d", &sym.val)
		sym.c = 0
		return
	}
	sym.c = c
}

// Abstract Syntax Tree functions

func node0(val int) *ast {
	return &ast{symbol{0, val}, nil}
}

func node2(c byte, x, y *ast) *ast {
	return &ast{symbol{c, 0}, []*ast{x, y}}
}

func eval(x *ast) int {
	switch x.c {
	case 0:
		return x.val
	case '+':
		return eval(x.args[0]) + eval(x.args[1])
	case '-':
		return eval(x.args[0]) - eval(x.args[1])
	case '*':
		return eval(x.args[0]) * eval(x.args[1])
	case '/':
		return eval(x.args[0]) / eval(x.args[1])
	}
	return 0
}

// Recursive Descent Parser functions

func factor() *ast {
	if sym.c == '(' {
		x := exp()
		if sym.c != ')' {
			panic("mismatched parentheses")
		}
		next()
		return x
	}
	next()
	return node0(sym.val)
}

func term() *ast {
	x := factor()
	for sym.c == '*' || sym.c == '/' {
		op := sym.c
		next()
		y := factor()
		switch op {
		case '*':
			x = node2('*', x, y)
		case '/':
			x = node2('/', x, y)
		}
	}
	return x
}

func exp() *ast {
	next()
	x := term()
	for sym.c == '+' || sym.c == '-' {
		op := sym.c
		next()
		y := term()
		switch op {
		case '+':
			x = node2('+', x, y)
		case '-':
			x = node2('-', x, y)
		}
	}
	return x
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	fmt.Println(eval(exp()))
}
