package doihaveabox

import (
	"api/v3"
	"app/services"
	"appengine"
	"appengine/urlfetch"
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/api/doihaveabox", doihaveabox)
}

func doihaveabox(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)

	var summonerName = r.URL.Query().Get("summonerName")

	summonerApi := v3.SummonerAPIImpl{HttpClient: client}
	var summonerService services.SummonerService
	summoner := summonerService.GetSummonerByName(summonerName, summonerApi)

	championMasteryApi := v3.ChampionMasteryAPIImpl{HttpClient: client}
	var championMasteryService services.ChampionMasteryService
	summonerChampionMasteryJsonStr := championMasteryService.GetJsonChampionMasteryBySummoner(summoner, championMasteryApi)

	fmt.Fprint(w, summonerChampionMasteryJsonStr)

}
