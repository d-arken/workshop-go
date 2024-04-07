package pokemon

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func GetPokemons() []Model {
	url := "https://pokeapi.co/api/v2/pokemon/?limit=10000"

	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var result struct {
		Results []Model `json:"results"`
	}
	err = json.Unmarshal(body, &result)

	return result.Results
}

type PokeAPIRes struct {
	Pokemon   *APIResponse
	TimeSpent string
}

func GetPokemon(url string) *PokeAPIRes {
	// Make the HTTP GET request
	t := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}
	spent := time.Since(t)

	var pokemon APIResponse
	err = json.Unmarshal(body, &pokemon)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return &PokeAPIRes{
		Pokemon:   &pokemon,
		TimeSpent: spent.String(),
	}
}
