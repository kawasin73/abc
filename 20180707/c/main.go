package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"math"
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
	n := float64(nextInt())
	m := float64(nextInt())
	d := float64(nextInt())
	fmt.Println((math.Pow(2, m-2)*(2*n-2*d*(d+1)) + 2*d*(d+1)) / math.Pow(n, m))
}
