package main

import "fmt"

func main() {

	var a int
	// Commit test
	fmt.Printf("Enter a number: ")
	fmt.Scanf("%v", &a)
	fmt.Println("Your numbers is:", a)

	/*for i := 0; i < a; i++ {
		if i%2 == 0 {
			fmt.Printf("%v is even\n", i)
		} else {
			fmt.Printf("%v is odd\n", i)
		}
	}*/
	if a%2 == 0 {
		fmt.Printf("%v is even\n", a)
	} else {
		fmt.Printf("%v is odd\n", a)
	}

}
