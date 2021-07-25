package domain

//go:generate mockgen -source=interfaces.go -destination=mocks.go -package=domain

type TimeConverter interface {
	ConvertTime(hours, minutes string) string
}

type UserInterface interface {
	GetTime() (string)
	ShowTime(string)
}