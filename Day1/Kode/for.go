package main

import "fmt"

func main() {
	
	//while
	i := 1
	for i<=3 {
		fmt.Println(i)
		i++
	}

	//for
	for j:=7;j<=9;j++{
		fmt.Println(j)
	}

	//repeat until
	for {
		fmt.Println("loop")
		break
	}
}