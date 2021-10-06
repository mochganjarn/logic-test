package leapyear

import (
	"bufio"
	"fmt"
	"os"
)

func Exetcute() {
	var first int
	fmt.Print("Masukan Tahun Awal: ")
	fmt.Scanf("%d", &first)
	var second int
	fmt.Print("Masukan Tahun Akhir: ")
	fmt.Scanf("%d", &second)
	for i := first; i <= second; i++ {
		if print_leapyear(i) {
			fmt.Printf("%v ", i)
		}
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\nEnter To Continue.....")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}

func print_leapyear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			}
		} else {
			return true
		}
	}
	return false
}
