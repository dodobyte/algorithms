package main

import (
	"bufio"
	"fmt"
	"os"
)

type prop struct {
	prec  int
	assoc string
}

var props = map[byte]prop{
	'+': prop{1, "left"},
	'-': prop{1, "left"},
	'*': prop{2, "left"},
	'/': prop{2, "left"},
}

type token struct {
	c   byte
	val int
}

var tok token
var out []token
var stack []token

var reader *bufio.Reader

func push(t token) {
	stack = append(stack, t)
}

func pop() token {
	t := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return t
}

func peek() token {
	return stack[len(stack)-1]
}

func next() {
	c, _ := reader.ReadByte()
	if c >= '0' && c <= '9' {
		reader.UnreadByte()
		fmt.Fscanf(reader, "%d", &tok.val)
		tok.c = 0
		return
	}
	tok.c = c
}

func shuntingYard() {
	for next(); tok.c != '\n'; next() {
		switch tok.c {
		case 0:
			out = append(out, tok)
		case '+', '-', '*', '/':
			pr := props[tok.c].prec
			as := props[tok.c].assoc
			for len(stack) > 0 {
				spr := props[peek().c].prec
				if pr < spr || (pr == spr && as == "left") {
					out = append(out, pop())
				} else {
					break
				}
			}
			push(tok)
		case '(':
			push(tok)
		case ')':
			lparn := false
			for len(stack) > 0 {
				t := pop()
				if t.c == '(' {
					lparn = true
					break
				}
				out = append(out, t)
			}
			if !lparn {
				panic("mismatched parentheses")
			}
		}
	}
	for len(stack) > 0 {
		t := pop()
		if t.c == '(' || t.c == ')' {
			panic("mismatched parentheses")
		}
		out = append(out, t)
	}
}

func eval() int {
	for _, t := range out {
		switch t.c {
		case '+':
			y, x := pop(), pop()
			x.val += y.val
			push(x)
		case '-':
			y, x := pop(), pop()
			x.val -= y.val
			push(x)
		case '*':
			y, x := pop(), pop()
			x.val *= y.val
			push(x)
		case '/':
			y, x := pop(), pop()
			x.val /= y.val
			push(x)
		case 0:
			push(t)
		}
	}
	return pop().val
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	for {
		shuntingYard()
		fmt.Printf("\t%d\n\n", eval())
		out = out[:0]
		stack = stack[:0]
	}
}
