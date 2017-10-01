package v3

import (
	"net/http"
	"src/api/util"
)

var riotSummonerApiURL = "https://na1.api.riotgames.com/lol/summoner/v3/summoners/"
var byName = "by-name/"

type SummonerAPIImpl struct {
	HttpClient *http.Client
}

func (champMasteryAPI SummonerAPIImpl) GetSummonerByName(summonerName string) (*util.RiotAPIHttpResponse, error) {
	var httpRequest util.RiotAPIHttpRequest
	summonerByNameRequestUrl := riotSummonerApiURL + byName + summonerName
	apiResponse, err := httpRequest.Do(summonerByNameRequestUrl, champMasteryAPI.HttpClient)

	return apiResponse, err
}
