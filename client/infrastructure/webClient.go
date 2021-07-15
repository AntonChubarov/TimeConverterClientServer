package infrastructure

import (
	"github.com/go-resty/resty/v2"
)

type WebClient struct {
	client *resty.Client
}

func NewWebClient() *WebClient {
	return &WebClient{
		resty.New(),
	}
}

func (w *WebClient) ConvertTimeOnServer(hours, minutes string) string {

	url := "http://localhost:8080/time?hours=" + hours + "&minutes=" + minutes
	resp, err:= w.client.R().Get(url)
	if err !=nil {
		panic(err)
	}
	if resp.IsSuccess() {
		return resp.String()
	} else {
		return resp.RawResponse.Status
	}
}