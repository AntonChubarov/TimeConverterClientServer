package infrastructure

import (
	"github.com/go-resty/resty/v2"
	"log"
)

type WebConverter struct {
	client *resty.Client
}

func NewWebConverter() *WebConverter {
	return &WebConverter{
		resty.New(),
	}
}

func (w *WebConverter) ConvertTime(hours, minutes string) string {
	url := "http://localhost:8080/time?hours=" + hours + "&minutes=" + minutes
	resp, err:= w.client.R().Get(url)
	if err !=nil {
		log.Println(err)
	}
	if resp.IsSuccess() {
		return resp.String()
	} else {
		return resp.RawResponse.Status
	}
}