package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"log"
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

var a = make([]int, 60)

func main() {
	sc.Split(bufio.ScanWords)

	for i := 2; i <= 30; i++ {
		a[i] = nextInt()
	}

	var kouho []int

	for i := 0; i < 256; i++ {
		x := i << 32
		min := 1<<32 - 1
		ok := true
		for j := 2; j <= 30; j ++ {
			k := x
			r := 0
			for k >= min {
				r += k % j
				k /= j
			}
			if r > a[j] {
				ok = false
				break
			}
		}
		if ok {
			kouho = append(kouho, x)
			log.Println("kouho1", x)
		}
	}

	var kouho2 []int

	for _, xx := range kouho {
		for i := 0; i < 256; i++ {
			x := i << 24
			x = xx + x
			min := 1<<24 - 1
			ok := true
			for j := 2; j <= 30; j ++ {
				k := x
				r := 0
				for k >= min {
					r += k % j
					k /= j
				}
				if r > a[j] {
					ok = false
					break
				}
			}
			if ok {
				kouho2 = append(kouho2, x)
			}
		}
	}

	var kouho3 []int

	for _, xx := range kouho2 {
		for i := 0; i < 256; i++ {
			x := i << 16
			x = xx + x
			min := 1<<16 - 1
			ok := true
			for j := 2; j <= 30; j ++ {
				k := x
				r := 0
				for k >= min {
					r += k % j
					k /= j
				}
				if r > a[j] {
					ok = false
					break
				}
			}
			if ok {
				kouho3 = append(kouho3, x)
			}
		}
	}

	var kouho4 []int

	for _, xx := range kouho3 {
		for i := 0; i < 256; i++ {
			x := i << 8
			x = xx + x
			min := 1<<8 - 1
			ok := true
			for j := 2; j <= 30; j ++ {
				k := x
				r := 0
				for k >= min {
					r += k % j
					k /= j
				}
				if r > a[j] {
					ok = false
					break
				}
			}
			if ok {
				kouho4 = append(kouho4, x)
			}
		}
	}

	for _, xx := range kouho4 {
		for i := 0; i < 256; i++ {
			x := xx + i

			if check(x) {
				fmt.Println(x)
				return
			}
		}
	}
	fmt.Println("invalid")
}

func check(test int) bool {
	for i := 2; i <= 30; i++ {
		if a[i] != sum(test, i) {
			return false
		}
	}
	return true
}

func sum(test int, i int) int {
	var s int
	for test > 0 {
		s += test % i
		test /= i
	}
	return s
}
