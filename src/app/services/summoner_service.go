package services

import (
	"encoding/json"
	"log"
	"src/api/interfaces"
	"src/app/dto"
)

type SummonerService struct {
}

func (summonerService *SummonerService) GetSummonerByName(summonerName string, summonerApi interfaces.SummonerAPI) *dto.Summoner {
	apiResponse, err := summonerApi.GetSummonerByName(summonerName)
	summoner := new(dto.Summoner)

	if err != nil {
		log.Printf("Error while querying Riot API: %s. Error: %s", apiResponse.Status, err)
		return summoner
	}

	jsonerr := json.Unmarshal(apiResponse.Data, summoner)

	if jsonerr != nil {
		log.Printf("Error while unmarshalling Riot API Response to SummonerChampionMastery: %s", jsonerr)
		return summoner
	}

	return summoner
}
