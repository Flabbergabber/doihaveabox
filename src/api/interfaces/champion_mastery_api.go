package interfaces

import (
	"github.com/Flabbergabber/doihaveabox/src/api/util"
)

type ChampionMasteryAPI interface {
	GetChampionMasteryBySummonerId(summonerId string) (*util.RiotAPIHttpResponse, error)
}
