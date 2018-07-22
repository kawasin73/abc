package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
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
	if (n < 1000) {
		fmt.Println("ABC")
	} else {
		fmt.Println("ABD")
	}
}
