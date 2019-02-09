package main

import (
	"bufio"
	"fmt"
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
	a := nextInt()
	b := nextInt()

	if a > b {
		fmt.Print(b, " ")
	} else {
		fmt.Print(a, " ")
	}

	min := (a + b) - n
	if min > 0 {
		fmt.Println(min)
	} else {
		fmt.Println(0)
	}
}
