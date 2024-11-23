package main

import "fmt"

func main() {
	d := []int{}
	for s := 2; s <= 26; s++ {
		alist := ListOfPrime(s, d)
		fmt.Printf("list of prime numbers of %d is : %v \n", s, alist)
	}

}

//This function takes an integer and returns true if that is a prime number, otherwise false.
func is_prime(n int) (b bool) {
	i := 2
	for i*i <= n {
		if n%i == 0 {
			return b
		}
		i++
	}
	return !b
}

//This function takes an integer and returns the smallest prime number that this integer is divisible by.
func firstPrime(n int) int {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			if is_prime(i) {
				return i
			}
		}
	}
	return n
}

//This function takes an integer and returns List of prime numbers whose product is equal to that integer.
func ListOfPrime(n int, t []int) []int {
	if n == 1 {
		return t
	}
	p := firstPrime(n)
	t = append(t, p)
	return ListOfPrime(n/p, t)
}
