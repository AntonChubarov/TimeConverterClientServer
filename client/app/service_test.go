package app

import (
	"client/domain"
	"github.com/golang/mock/gomock"
	"testing"
)

type Case struct {
	desc string
	converter domain.TimeConverter
	ui domain.UserInterface
}

func TestService(t *testing.T)  {
	ctrl := gomock.NewController(t)

	inputTime := "18:08"
	inputHours := "18"
	inputMinutes := "08"
	outputTime := "6:08 PM"

	cases := []Case {
		{
			desc:"service when starts should call all methods at least once",
			converter: func(ctrl *gomock.Controller) *domain.MockTimeConverter {
				mock:=domain.NewMockTimeConverter(ctrl)
				mock.EXPECT().ConvertTime(inputHours, inputMinutes).Return(outputTime).MinTimes(1)
				return mock
			}(ctrl),
			ui: func(ctrl *gomock.Controller) *domain.MockUserInterface {
				mock:=domain.NewMockUserInterface(ctrl)
				mock.EXPECT().GetTime().Return(inputTime).MinTimes(1)
				mock.EXPECT().ShowTime(outputTime).MinTimes(1)
				return mock
			}(ctrl),
		},
	}

	for i := range cases {
		t.Run(cases[i].desc, func(t *testing.T) {
			RunWithoutLoop(cases[i].converter, cases[i].ui)
		})
	}
}