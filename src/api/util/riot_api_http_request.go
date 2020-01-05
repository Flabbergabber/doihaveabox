package util

import (
	"github.com/Flabbergabber/doihaveabox/src/api/config"

	"bytes"
	"net/http"
)

type RiotAPIHttpRequest struct {
	RequestURL string
}

func (riotApiRequest *RiotAPIHttpRequest) Do(requestURL string) (*RiotAPIHttpResponse, error) {
	apiResponse := new(RiotAPIHttpResponse)

	req, err := http.NewRequest("GET", requestURL, nil)

	if err != nil {
		return apiResponse, err
	}

	apiKey := config.GetInstance().GetRiotApiKey()
	req.Header.Set("X-Riot-Token", string(apiKey))

	client := &http.Client{}
	httpResponse, err := client.Do(req)

	var buf bytes.Buffer
	buf.ReadFrom(httpResponse.Body)
	apiResponse.Data = buf.Bytes()
	apiResponse.Status = httpResponse.Status

	if err != nil {
		return apiResponse, err
	} else {
		return apiResponse, nil
	}
}
