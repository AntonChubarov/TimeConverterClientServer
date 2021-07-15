package app

import (
	"client/infrastructure"
	"log"
)

func Run() {
	client := infrastructure.NewWebClient()

	for {
		time := infrastructure.GetTime()
		if hours, minutes, err := infrastructure.TryToParse(time); err == nil {
			timeFromServer := client.ConvertTimeOnServer(hours, minutes)
			infrastructure.ShowTime(timeFromServer)
		} else {
			log.Println(err)
			continue
		}
	}
}