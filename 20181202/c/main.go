package main

import (
	"bufio"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextString() string {
	if !sc.Scan() {
		panic("failed scan string")
	}
	return sc.Text()
}

func nextInt() int {
	if !sc.Scan() {
		panic("failed scan int")
	}
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	ns := make([]int, 10)

	for n != 0 {
		ns = append(ns, n%10)
		n /= 10
	}

	for len(ns) > 0 {

	}
}

func sum(ns []int) int {
	if len(ns) == 0 {
		return 0
	}
	switch ns[0] {
	case 0, 1, 2:
		return 0
	case 3:

	case 4:
	case 5:
	case 6:
	case 7:
		sum(ns[1:])
	case 8, 9:
		return full(len(ns))
	}
}

func full(n int) int {
	return pow(3, n) - (3 * pow(2, n-1)) + 3
}

func pow(i, n int) int {
	result := 1
	for k := 0; k < n; k++ {
		result *= i
	}
	return result
}

func have7(ns []int) int {

}
