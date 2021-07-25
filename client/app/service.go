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

// RunWithoutLoop func is for testing only
func RunWithoutLoop (converter domain.TimeConverter, ui domain.UserInterface) {
	// this code should be the copy of code of Run() func that inside the for{} loop, and be without a continue operator
	time := ui.GetTime()
	if hours, minutes, err := TryToParse(time); err == nil {
		timeFromServer := converter.ConvertTime(hours, minutes)
		ui.ShowTime(timeFromServer)
	} else {
		log.Println(err)
	}
}