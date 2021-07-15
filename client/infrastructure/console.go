package infrastructure

import "fmt"

func GetTime() (input string) {
	fmt.Println("Type the time in EU format (example: 15:36) ant press enter")
	fmt.Scan(&input)
	return
}

func ShowTime(time string) {
	fmt.Println("Your time in US format:")
	fmt.Println(time)
}
