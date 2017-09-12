package main

import (
	"fmt"
)

// 100 unsorted random number between 1-1000
var list = []int{
	801, 468, 485, 216, 569, 388, 875, 984, 302, 815,
	793, 704, 632, 559, 363, 784, 846, 647, 800, 282,
	833, 816, 858, 667, 712, 823, 790, 591, 281, 739,
	176, 564, 436, 335, 214, 409, 399, 379, 849, 925,
	387, 24, 611, 661, 926, 286, 572, 317, 333, 498,
	34, 437, 202, 7, 459, 58, 536, 779, 964, 697,
	12, 434, 982, 513, 127, 93, 299, 890, 912, 212,
	676, 821, 321, 232, 233, 736, 979, 837, 920, 270,
	590, 118, 162, 829, 1, 990, 929, 162, 909, 72,
	813, 766, 123, 514, 980, 343, 233, 74, 144, 340,
}

// Quicksort algorithm that uses Lomuto partition scheme
func qsort(l []int) {
	n := len(l)
	if n < 2 {
		return
	}
	pos := 0
	piv := &l[n-1]
	for i := range l {
		if l[i] < *piv {
			l[i], l[pos] = l[pos], l[i]
			pos++
		}
	}
	l[pos], *piv = *piv, l[pos]
	qsort(l[:pos])
	qsort(l[pos+1:])
}

// Bubble sort algorithm for verification
func bubble(l []int) {
	n := len(list)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if l[j] < l[i] {
				l[i], l[j] = l[j], l[i]
			}
		}
	}
}

func main() {
	qsort(list)
	fmt.Println(list)
}
