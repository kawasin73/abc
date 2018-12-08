package main

import (
	"fmt"
)

var (
	memoTotal  []uint64
	memoMiddle []uint64
	memoP      []uint64
)

func main() {
	var n, x uint64
	fmt.Scanf("%d %d", &n, &x)

	memoTotal = make([]uint64, n+1)
	memoMiddle = make([]uint64, n+1)
	memoP = make([]uint64, n+1)
	memoTotal[0] = 1
	memoP[0] = 1
	for i := uint64(1); i <= n; i++ {
		memoTotal[i] = memoTotal[i-1]*2 + 3
		memoMiddle[i] = 2 + memoTotal[i-1]
		memoP[i] = memoP[i-1]*2 + 1
	}
	fmt.Println(search(n, x))
}

func search(n uint64, rest uint64) uint64 {
	if rest == 0 {
		return 0
	} else if rest == 1 {
		if n == 0 {
			return 1
		}
		return 0
	}

	if rest == memoTotal[n] {
		return memoP[n]
	} else if rest > memoMiddle[n] {
		return search(n-1, rest-memoMiddle[n]) + memoP[n-1] + 1
	} else if rest == memoMiddle[n] {
		return memoP[n-1] + 1
	} else {
		return search(n-1, rest-1)
	}
}
