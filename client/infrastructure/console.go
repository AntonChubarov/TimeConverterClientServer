package infrastructure

import "fmt"

type Console struct {}

func NewConsole() *Console {
	return &Console{}
}

func (c *Console) GetTime() (input string) {
	fmt.Println("Type the time in EU format (example: 15:36) ant press enter")
	fmt.Scan(&input)
	return
}

func (c *Console) ShowTime(time string) {
	fmt.Println("Your time in US format:")
	fmt.Println(time)
}
