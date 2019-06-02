package main

import (
	"fmt"
	"math"
)

func Buddy(start, limit int) []int {
	for m := start; m < limit; m++ {
		msum := divisor(m)
		n := msum - 1
		if n < start {
			continue
		}
		nsum := divisor(n)
		if nsum == m+1 {
			if m < n {
				return []int{m, n}
			} else {
				return []int{n, m}
			}
		}
	}
	return nil
}

func divisor(n int) int {
	temp := 1
	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 {
			if n/i == i {
				temp += i
			} else {
				temp += i
				temp += n / i
			}

		}
	}
	return temp
}

func main() {
	fmt.Println(Buddy(10, 50))
}
