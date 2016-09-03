package main

import "fmt"

func plus(a int, b int) int {
	return a+b
}

func plusPlus(a, b, c int) int {
	return a+b+c
}

func hitungNPertama(n int) int {
	var counter int
	for i := 1; i <= n; i++ {
		counter+=i
	}
	return counter
}

func main() {
	res := plus(1,2)
	fmt.Println("1+2 = ",res)

	res = plusPlus(1,2,3) 
	fmt.Println("1+2+3 = ", res)

	fmt.Println("10 suku pertama jumlahnya ",hitungNPertama(10))
}