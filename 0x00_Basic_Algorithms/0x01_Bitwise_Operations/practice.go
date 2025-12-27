package main

import "fmt"

/**
 * calculate a^b mod p 1<= a,b,p <= 10^9
 */
func fastPower(a, b, p int64) int64 {
	ans := int64(1) % p;
	for ; b > 0; b >>= 1 {
		if b & 1 == 1 {
			ans = (ans * a) % p;
		}
		a = (a * a) % p;
	}
	return ans
}

func main() {
    fmt.Println(fastPower(3, 5, 100)) 
}