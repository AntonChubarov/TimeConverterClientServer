package app

import (
	"client/domain"
	"log"
)

func Run(converter domain.TimeConverter, ui domain.UserInterface) {
	for {
		time := ui.GetTime()
		if hours, minutes, err := TryToParse(time); err == nil {
			timeFromServer := converter.ConvertTime(hours, minutes)
			ui.ShowTime(timeFromServer)
		} else {
			log.Println(err)
			continue
		}
	}
}