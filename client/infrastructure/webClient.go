package infrastructure

import (
	"github.com/go-resty/resty/v2"
	"strconv"
)

type WebClient struct {
	client *resty.Client
}

func NewWebClient() *WebClient {
	return &WebClient{
		resty.New(),
	}
}

func (w *WebClient) ConvertTimeOnServer(hours, minutes int) string {

	url := "http://localhost:8080/time?hours=" + strconv.Itoa(hours) + "&minutes=" + strconv.Itoa(minutes)
	resp, err:= w.client.R().Get(url)
	if err !=nil {
		panic(err)
	}
	return resp.String()
}