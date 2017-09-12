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

var sym symbol
var reader *bufio.Reader

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

func factor() {
	if sym.c == '(' {
		exp()
		if sym.c != ')' {
			panic("mismatched parentheses")
		}
		next()
		return
	}
	fmt.Printf("%d ", sym.val)
	next()
}

func term() {
	factor()
	for sym.c == '*' || sym.c == '/' {
		op := sym.c
		next()
		factor()
		fmt.Printf("%c ", op)
	}
}

func exp() {
	next()
	term()
	for sym.c == '+' || sym.c == '-' {
		op := sym.c
		next()
		term()
		fmt.Printf("%c ", op)
	}
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	exp()
}
