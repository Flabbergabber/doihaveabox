package interfaces

import (
	"api/util"
)

type ChampionMasteryAPI interface {
	GetChampionMasteryBySummonerId(summonerId int) (*util.RiotAPIHttpResponse, error)
}
