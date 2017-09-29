package hello

import (
    "fmt"
    "net/http"
    "log"
    "appengine"
    "appengine/urlfetch"
    "bytes"
    "encoding/json"
    //"io/ioutil"
    "strconv"
)

type Summoner struct {
    Id int `json:"id"`
    AccountId int `json:"accountId"`
    Name string `json:"name"`
    ProfileIconId int `json:"profileIconId"`
    RevisionDate int `json:"revisionDate"`
    SummonerLevel int `json:"summonerLevel"`
}

type ChampionMastery struct {
    ChampionLevel int `json:"championLevel"`
    ChestGranted bool `json:"chestGranted"`
    ChampionPoints int `json:"championPoints"`
    ChampionPointsSinceLastLevel int `json:"championPointsSinceLastLevel"`
    PlayerId int `json:"playerId"`
    ChampionPointsUntilNextLevel int `json:"championPointsUntilNextLevel"`
    TokensEarned int `json:"tokensEarned"`
    ChampionId int `json:"championId"`
    LastPlayTime int `json:"lastPlayTime"`
}

type SummonerMasteryResult []ChampionMastery

func init() {
    http.HandleFunc("/api/doihaveabox", doihaveabox)
}

func doihaveabox(w http.ResponseWriter, r *http.Request) {
    ctx := appengine.NewContext(r)
    client := urlfetch.Client(ctx)
    var riotSummonerApiURL = "https://na1.api.riotgames.com/lol/summoner/v3/summoners/by-name/"
    var riotMasteryApiURL = "https://na1.api.riotgames.com/lol/champion-mastery/v3/champion-masteries/by-summoner/"
    var apiKey = "RGAPI-219fdb89-b870-4ac8-8716-3f1913e185e7"

    var summonerName = r.URL.Query().Get("summonerName")

    if len(summonerName) == 0 {
        fmt.Fprint(w, "Please enter a Summoner Name")
        return
    }

    response, err := client.Get(riotSummonerApiURL + summonerName + "?api_key=" + apiKey)
    log.Print("Response: " + response.Status)
    if err != nil {
        fmt.Print(w, err.Error())
        log.Print(err.Error())
    } else {
        buf := new(bytes.Buffer)
        buf.ReadFrom(response.Body)
        responseStr := buf.String()

        j := []byte(responseStr)

        var summoner Summoner
        jsonerr := json.Unmarshal(j, &summoner)

        if jsonerr != nil {
            log.Print(jsonerr)
        }

        response, err := client.Get(riotMasteryApiURL + strconv.Itoa(summoner.Id) + "?api_key=" + apiKey)
        buf = new(bytes.Buffer)
        buf.ReadFrom(response.Body)
        responseStr = buf.String()

        j = []byte(responseStr)

        var summonerMasteryResult SummonerMasteryResult
        jsonerr = json.Unmarshal(j, &summonerMasteryResult)

        log.Print(summonerMasteryResult[0].ChestGranted)

        if err != nil {
            log.Fatal(err)
        }
        fmt.Fprint(w, responseStr)
    }
}
