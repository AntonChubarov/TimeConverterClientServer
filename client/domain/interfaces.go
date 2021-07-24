package domain

type TimeConverter interface {
	ConvertTime(hours, minutes string) string
}

type UserInterface interface {
	GetTime() (string)
	ShowTime(string)
}