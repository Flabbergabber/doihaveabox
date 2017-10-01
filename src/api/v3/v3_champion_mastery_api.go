package v3

import (
	"net/http"
	"src/api/util"
	"strconv"
)

var riotMasteryApiURL = "https://na1.api.riotgames.com/lol/champion-mastery/v3/champion-masteries/"
var bySummoner = "by-summoner/"

type ChampionMasteryAPIImpl struct {
	HttpClient *http.Client
}

func (champMasteryAPI ChampionMasteryAPIImpl) GetChampionMasteryBySummonerId(summonerId int) (*util.RiotAPIHttpResponse, error) {
	var httpRequest util.RiotAPIHttpRequest
	championMasteryBySummonerIdUrl := riotMasteryApiURL + bySummoner + strconv.Itoa(summonerId)
	apiResponse, err := httpRequest.Do(championMasteryBySummonerIdUrl, champMasteryAPI.HttpClient)

	return apiResponse, err
}
