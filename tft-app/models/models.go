package models

// LeagueListDTO represents the response body of the Riot API for the league endpoint
type LeagueListDTO struct {
	LeagueId string          `json:"leagueId"`
	Entries  []LeagueItemDTO `json:"entries"`
	Tier     string          `json:"tier"`
	Name     string          `json:"name"`
	Queue    string          `json:"queue"`
}

// LeagueItemDTO represents the individual player data in a league
type LeagueItemDTO struct {
	FreshBlood   bool   `json:"freshBlood"`
	Wins         int    `json:"wins"`
	Inactive     bool   `json:"inactive"`
	Veteran      bool   `json:"veteran"`
	HotStreak    bool   `json:"hotStreak"`
	Rank         string `json:"rank"`
	LeaguePoints int    `json:"leaguePoints"`
	Losses       int    `json:"losses"`
	SummonerId   string `json:"summonerId"`
	Puuid        string `json:"puuid"`
}

// LeagueItemWithName is an extended version of LeagueItemDTO that includes the SummonerName
type LeagueItemWithName struct {
	LeagueItemDTO
	SummonerName string `json:"summonerName"`
}

// AccountDto represents the account data that includes the game name and tag line for a summoner
type AccountDto struct {
	Puuid    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}
