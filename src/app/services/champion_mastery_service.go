package services

import (
	"src/app/dto"
	"src/api/interfaces"
	"encoding/json"
	"log"
)

type ChampionMasteryService struct {
}

func (champMasteryService *ChampionMasteryService) GetJsonChampionMasteryBySummoner(summonerDto *dto.Summoner, masteryApi interfaces.ChampionMasteryAPI) string {
	apiResponse, err := masteryApi.GetChampionMasteryBySummonerId(summonerDto.Id)
	summonerChampionMastery := new(dto.SummonerChampionMastery)

	if err != nil {
		log.Printf("Error while querying Riot API: %s. Error: %s", apiResponse.Status, err)
		return ""
	}

	jsonerr := json.Unmarshal(apiResponse.Data, summonerChampionMastery)

	if jsonerr != nil {
		log.Printf("Error while unmarshalling Riot API Response to SummonerChampionMastery: %s", jsonerr)
		return ""
	}

	responseStr := string(apiResponse.Data)

	return responseStr
}