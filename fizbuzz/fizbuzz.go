package fizbuzz

import (
	"fmt"
	"strconv"
)

func Execute() {
	var i int
	fmt.Print("Enter count: ")
	fmt.Scanf("%d", &i)
	var result []string
	for j := 1; j <= i; j++ {
		if j%3 == 0 {
			if j%5 == 0 {
				result = append(result, "FizzBuzz")
			} else {
				result = append(result, "Fizz")
			}
		} else if j%5 == 0 {
			result = append(result, "Buzz")
		} else {
			res := strconv.Itoa(j)
			result = append(result, res)
		}
	}
	fmt.Println(result)
}
