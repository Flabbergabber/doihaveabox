package v4

import (
	"github.com/Flabbergabber/doihaveabox/src/api/util"
)

var riotMasteryApiURL = "https://na1.api.riotgames.com/lol/champion-mastery/v4/champion-masteries/"
var bySummoner = "by-summoner/"

type ChampionMasteryAPIImpl struct {
}

func (champMasteryAPI ChampionMasteryAPIImpl) GetChampionMasteryBySummonerId(summonerId string) (*util.RiotAPIHttpResponse, error) {
	var httpRequest util.RiotAPIHttpRequest
	championMasteryBySummonerIdUrl := riotMasteryApiURL + bySummoner + summonerId
	apiResponse, err := httpRequest.Do(championMasteryBySummonerIdUrl)

	return apiResponse, err
}
