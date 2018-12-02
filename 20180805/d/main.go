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

const D = 1000000007

func main() {
	sc.Split(bufio.ScanWords)

	s := nextString()

	a := 0
	b := 0
	c := 0

	totalQ := 0

	for _, c := range s {
		if c == '?' {
			totalQ++
		}
	}

	for _, ca := range s {
		switch ca {
		case 'A':
			a ++
			a %= D
		case 'B':
			b += a
			b %= D
		case 'C':
			c += b * pow(totalQ)
			c %= D
		case '?':
			totalQ--
			c += b * pow(totalQ)
			b *= 3
			b += a
			a ++
			a %= D
			b %= D
			c %= D
		}
		fmt.Println(a, b, c)
	}

	fmt.Println(c)
}

var mem = make(map[int]int)

func pow(n int) int {
	if result, ok := mem[n]; ok {
		return result
	}
	result := 1
	for i := 0; i < n; i++ {
		result *= 3
		result %= D
	}
	mem[n] = result
	return result
}
