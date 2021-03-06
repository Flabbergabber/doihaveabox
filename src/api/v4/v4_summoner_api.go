package v4

import (
	"github.com/Flabbergabber/doihaveabox/src/api/util"
)

var riotSummonerApiURL = "https://na1.api.riotgames.com/lol/summoner/v4/summoners/"
var byName = "by-name/"

type SummonerAPIImpl struct {
}

func (champMasteryAPI SummonerAPIImpl) GetSummonerByName(summonerName string) (*util.RiotAPIHttpResponse, error) {
	var httpRequest util.RiotAPIHttpRequest
	summonerByNameRequestUrl := riotSummonerApiURL + byName + summonerName
	apiResponse, err := httpRequest.Do(summonerByNameRequestUrl)

	return apiResponse, err
}
