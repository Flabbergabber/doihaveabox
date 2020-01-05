package dto

type ChampionMastery struct {
	ChampionLevel                int    `json:"championLevel"`
	ChestGranted                 bool   `json:"chestGranted"`
	ChampionPoints               int    `json:"championPoints"`
	ChampionPointsSinceLastLevel int    `json:"championPointsSinceLastLevel"`
	SummonnerId                  string `json:"summonerId"`
	ChampionPointsUntilNextLevel int    `json:"championPointsUntilNextLevel"`
	TokensEarned                 int    `json:"tokensEarned"`
	ChampionId                   int    `json:"championId"`
	LastPlayTime                 int    `json:"lastPlayTime"`
}

type SummonerChampionMastery []ChampionMastery
