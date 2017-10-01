package interfaces

import (
	"src/api/util"
)

type ChampionMasteryAPI interface {
	GetChampionMasteryBySummonerId(summonerId int) (*util.RiotAPIHttpResponse, error)
}
