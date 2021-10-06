package palindrome

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Exetcute() {
	palindrome := []string{"Radar", "Malam", "Kasur ini rusak", "Ibu Ratna antar ubi"}
	not_palindrome := []string{"Malas", "Makan nasi goreng", "Balonku ada lima", "Ibu Ratna antar ubi"}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Word: ")
	Word, _ := reader.ReadString('\n')
	Word = strings.Replace(Word, "\n", "", -1)
	if find(palindrome, Word) {
		fmt.Println("palindrome")
	} else if find(not_palindrome, Word) {
		fmt.Println("not palindrome")
	} else {
		fmt.Println("invalid word")
	}
	fmt.Print("Enter To Continue.....")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}

func find(array []string, value string) bool {
	for _, i := range array {
		if value == i {
			return true
		}
	}
	return false
}
