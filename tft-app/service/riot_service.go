package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"tft-app/models"
)

const (
	// API endpoints
	challengerURL  = "https://euw1.api.riotgames.com/tft/league/v1/challenger"
	grandmasterURL = "https://euw1.api.riotgames.com/tft/league/v1/grandmaster"
	masterURL      = "https://euw1.api.riotgames.com/tft/league/v1/master"
	// Riot API key (ensure you set this up as an env variable)
	apiKey = "RGAPI-4521fb89-ac56-4a00-b16c-2028deadb489"
	// Default page size
	pageSize = 25
)

func FetchLeagueData(url string) ([]models.LeagueItemWithName, error) {
	// Construct URL with pagination parameters
	fullURL := fmt.Sprintf("%s?queue=RANKED_TFT&api_key=%s", url, apiKey)

	// Make the API call
	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %v", err)
	}
	defer resp.Body.Close()

	// Check for rate limit error
	if resp.StatusCode == 429 {
		return nil, fmt.Errorf("rate limit exceeded")
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Check if the API returned a non-200 status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned non-OK status: %s", resp.Status)
	}

	// Parse the response into LeagueListDTO struct
	var leagueData models.LeagueListDTO
	err = json.Unmarshal(body, &leagueData)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	// Prepare a list of league entries with summoner names
	var result []models.LeagueItemWithName

	for i, entry := range leagueData.Entries {
		if i >= pageSize {
			break
		}
		// Fetch the summoner name using the puuid
		summonerName, err := GetSummonerName(entry.Puuid)
		if err != nil {
			// Handle error gracefully, continue with just the ID if name can't be fetched
			summonerName = "Unknown"
		}

		// Add the name to the entry
		leagueEntry := models.LeagueItemWithName{
			LeagueItemDTO: entry,
			SummonerName:  summonerName,
		}

		result = append(result, leagueEntry)
	}

	// Return the list of league entries with names
	return result, nil
}

// GetSummonerName - Fetches the summoner name by puuid
func GetSummonerName(puuid string) (string, error) {
	// Construct the URL to get account info by puuid
	url := fmt.Sprintf("https://europe.api.riotgames.com/riot/account/v1/accounts/by-puuid/%s?api_key=%s", puuid, apiKey)

	// Make the HTTP request to Riot API
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to make request to Riot API for summoner name: %v", err)
	}
	defer resp.Body.Close()

	// Check for rate limit error
	if resp.StatusCode == 429 {
		return "", fmt.Errorf("rate limit exceeded")
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body for summoner name: %v", err)
	}

	// Check if the API returned a non-200 status code
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API returned non-OK status: %s", resp.Status)
	}

	// Parse the response to get the summoner's name
	var accountData models.AccountDto
	err = json.Unmarshal(body, &accountData)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal response for summoner name: %v", err)
	}

	// Return the summoner name
	return accountData.GameName + "#" + accountData.TagLine, nil
}

// FetchAllChallengers retrieves all challengers using pagination
func FetchAllChallengers() ([]models.LeagueItemWithName, error) {
	return FetchLeagueData(challengerURL)
}

// FetchAllGrandmasters retrieves all grandmasters using pagination
func FetchAllGrandmasters() ([]models.LeagueItemWithName, error) {
	return FetchLeagueData(grandmasterURL)
}

// FetchAllMasters retrieves all masters using pagination
func FetchAllMasters() ([]models.LeagueItemWithName, error) {
	return FetchLeagueData(masterURL)
}
