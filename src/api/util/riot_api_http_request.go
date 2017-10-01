package util

import (
	"net/http"
	"src/api/config"
	"bytes"
)

type RiotAPIHttpRequest struct {
	RequestURL string
	HttpClient *http.Client
}

func (riotApiRequest *RiotAPIHttpRequest) Do(requestURL string, httpClient *http.Client) (*RiotAPIHttpResponse, error) {
	apiResponse := new(RiotAPIHttpResponse)

	req, err := http.NewRequest("GET", requestURL, nil)

	if err != nil {
		return apiResponse, err
	}

	var apiKey = config.GetInstance().GetRiotApiKey()
	req.Header.Set("X-Riot-Token", string(apiKey))
	httpResponse, err := httpClient.Do(req)

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