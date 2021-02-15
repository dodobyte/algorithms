# Algorithms in Go

Some stand-alone algorithms implemented in Go. See my [gist](https://gist.github.com/dodobyte) for the C version of some of the algorithms.

- [Sorting Algorithms](https://en.wikipedia.org/wiki/Sorting_algorithm)
  - `sort.go` [Quicksort](https://en.wikipedia.org/wiki/Quicksort) and Bubble sort algorithms.
- [Recursive Descent Parser](https://en.wikipedia.org/wiki/Recursive_descent_parser)
  - `rdp_rpn.go` Convert a mathematical expression into [Reverse Polish Notation](https://en.wikipedia.org/wiki/Reverse_Polish_notation).
  - `rdp_expr.go` Evaluate mathematical expressions and print results.
- [Abstract Syntax Tree](https://en.wikipedia.org/wiki/Abstract_syntax_tree)
  - `ast.go` Parse mathematical expressions to build an Abstract Syntax Tree. Evaluate the final AST.
- [Trie Data Structure](https://en.wikipedia.org/wiki/Trie)
  - `trie.go` Trie implementation. insert & lookup implemented with some test code.
- [Shunting Yard Algorithm](https://en.wikipedia.org/wiki/Shunting-yard_algorithm)
  - `shunting_yard.go` Parse mathematical expressions with shunting-yard algorithm and evaluate with a [stack machine](https://en.wikipedia.org/wiki/Stack_machine).
- [Ray casting](https://en.wikipedia.org/wiki/Ray_casting)
  - See [dodobyte/ray-casting](https://github.com/dodobyte/ray-casting)
- [A* pathfinding](https://en.wikipedia.org/wiki/A*_search_algorithm)
  - See [dodobyte/astar](https://github.com/dodobyte/astar)
