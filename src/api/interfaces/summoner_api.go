package interfaces

import (
	"github.com/Flabbergabber/doihaveabox/src/api/util"
)

type SummonerAPI interface {
	GetSummonerByName(summonerName string) (*util.RiotAPIHttpResponse, error)
}
