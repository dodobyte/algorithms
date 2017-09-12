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

func factor() int {
	if sym.c == '(' {
		v := exp()
		if sym.c != ')' {
			panic("mismatched parentheses")
		}
		next()
		return v
	}
	next()
	return sym.val
}

func term() int {
	v1 := factor()
	for sym.c == '*' || sym.c == '/' {
		op := sym.c
		next()
		v2 := factor()
		switch op {
		case '*':
			v1 *= v2
		case '/':
			v1 /= v2
		}
	}
	return v1
}

func exp() int {
	next()
	v1 := term()
	for sym.c == '+' || sym.c == '-' {
		op := sym.c
		next()
		v2 := term()
		switch op {
		case '+':
			v1 += v2
		case '-':
			v1 -= v2
		}
	}
	return v1
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	fmt.Println(exp())
}
