package nearfib

import "fmt"

func Execute(arr []int) string {
	var count int
	for _, v := range arr {
		count += v
	}
	current := 0
	first, second := 0, 1
	for current < count {
		current = first
		first, second = second, first+second
	}
	bot := (count - current) * -1
	up := (count - first) * -1
	if bot > up {
		return fmt.Sprintf("Nearest fibonaci adalah %v", up)
	} else {
		return fmt.Sprintf("Nearest fibonaci adalah %v", bot)
	}
}
