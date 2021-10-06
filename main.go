package main

import (
	"fmt"

	"github.com/mochganjarn/logic-test/fizbuzz"
	"github.com/mochganjarn/logic-test/leapyear"
	"github.com/mochganjarn/logic-test/nearfib"
	"github.com/mochganjarn/logic-test/palindrome"
	"github.com/mochganjarn/logic-test/reverse"
)

func main() {
	var i int
	for i != 6 {
		fmt.Println("Menu")
		fmt.Println("1. Palindrome")
		fmt.Println("2. Leap Year")
		fmt.Println("3. Reverse Word")
		fmt.Println("4. Nearest Fibonaci")
		fmt.Println("5. FizBuzz")
		fmt.Println("6. Exit")

		fmt.Print("Enter Number: ")
		fmt.Scanf("%d", &i)
		switch i {
		case 1:
			palindrome.Exetcute()
		case 2:
			leapyear.Exetcute()
		case 3:
			reverse.Exetcute()
		case 4:
			fmt.Println(nearfib.Execute([]int{1, 2, 100}))
			fmt.Print("Enter to continue...")
			fmt.Scanf("%d", &i)
		case 5:
			fizbuzz.Execute()
		}
	}
}
