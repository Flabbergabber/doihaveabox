package interfaces

import (
	"src/api/util"
)

type SummonerAPI interface {
	GetSummonerByName(summonerName string) (*util.RiotAPIHttpResponse, error)
}