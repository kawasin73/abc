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
	as := make([]uint64, n)
	for i := 0; i < n; i++ {
		as[i] = uint64(nextInt())
	}

	var r int
	var l int

	var sum uint64
	var count int

	for l = 0; l < n; l++ {
		for as[l]&sum != 0 {
			sum -= as[r]
			r += 1
		}
		//fmt.Printf("r: %v, l: %v\n", r, l)
		sum += as[l]
		count += l - r + 1
	}

	fmt.Println(count)
}

func check(as []uint64) bool {
	if len(as) == 0 {
		return false
	}
	xor := as[0]
	sum := as[0]
	for _, a := range as[1:] {
		xor ^= a
		sum += a
		if xor != sum {
			return false
		}
	}
	//if xor == sum {
	//	fmt.Printf("xor: %v, sum: %v\n", xor, sum)
	//}
	return xor == sum
}
