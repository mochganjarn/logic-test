package reverse

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Exetcute() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Word: ")
	Word, _ := reader.ReadString('\n')
	Word = strings.Replace(Word, "\n", "", -1)

	res := strings.Split(Word, " ")
	var result string
	var reversed string
	for index, i := range res {
		for j := len(i) - 1; j >= 0; j-- {
			reversed += fmt.Sprintf("%c", i[j])
		}
		if index == len(res)-1 {
			result += reversed
			reversed = ""
		} else {
			result += reversed + " "
			reversed = ""
		}
	}
	fmt.Println(result)
	fmt.Print("Enter To Continue.....")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
