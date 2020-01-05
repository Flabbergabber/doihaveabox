package main

import (
	"fmt"
	"log"
	"os"

	v4 "github.com/Flabbergabber/doihaveabox/src/api/v4"
	"github.com/Flabbergabber/doihaveabox/src/app/services"

	"net/http"
)

func main() {
	http.HandleFunc("/api/doihaveabox", doihaveabox)

	index := http.FileServer(http.Dir("www/app/"))
	http.Handle("/", index)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func doihaveabox(w http.ResponseWriter, r *http.Request) {
	var summonerName = r.URL.Query().Get("summonerName")

	summonerApi := v4.SummonerAPIImpl{}
	var summonerService services.SummonerService
	summoner := summonerService.GetSummonerByName(summonerName, summonerApi)

	championMasteryApi := v4.ChampionMasteryAPIImpl{}
	var championMasteryService services.ChampionMasteryService
	summonerChampionMasteryJsonStr := championMasteryService.GetJsonChampionMasteryBySummoner(summoner, championMasteryApi)

	fmt.Fprint(w, summonerChampionMasteryJsonStr)
}
