package main

import (
	"fmt"
	"os"
)

// --- Trie implementation ---

type trie struct {
	value interface{}
	word  bool
	child [256]*trie
}

func (root *trie) insert(key string, value interface{}) {
	p := root
	for _, k := range key {
		if p.child[k] == nil {
			p.child[k] = new(trie)
		}
		p = p.child[k]
	}
	p.value, p.word = value, true
}

func (root *trie) lookup(key string) (interface{}, bool) {
	p := root
	for _, k := range key {
		if p.child[k] == nil {
			return nil, false
		}
		p = p.child[k]
	}
	return p.value, p.word
}

// --- Test codes ---

var testData = map[string]int{
	"dodo":      10,
	"alena":     100,
	"wirth":     1000,
	"hoare":     1100,
	"ritchie":   1200,
	"dijkstra":  1300,
	"thompson":  1400,
	"kernighan": 900,
}

func test(root *trie) {
	for key, val := range testData {
		trval, ok := root.lookup(key)
		if !ok || val != trval {
			fmt.Fprintf(os.Stderr, "FAIL %s:%d\n\n", key, val)
			return
		}
	}
	fmt.Fprintf(os.Stderr, "\nPASS\n\n")
}

func main() {
	root := new(trie)

	for key, val := range testData {
		root.insert(key, val)
	}

	test(root)            // Should pass
	testData["bad"] = 666 // Create a failure case
	test(root)            // Should fail

	// Some individual test cases
	fmt.Println(root.lookup("dod"))
	fmt.Println(root.lookup("dodo"))
	fmt.Println(root.lookup("dodobig"))
	fmt.Println(root.lookup("a"))
	fmt.Println(root.lookup("alen"))
	fmt.Println(root.lookup("alena"))
	fmt.Println(root.lookup("alena1"))
}
