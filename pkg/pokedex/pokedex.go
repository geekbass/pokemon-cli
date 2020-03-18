package pokedex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Pokedex struct {
	Descriptions []PokedexDescriptions `json:"descriptions"`
	ID           string                `json:"id"`
	Name         string                `json:"name"`
	Pokemon      []PokemonEntries      `json:"pokemon_entries"`
}

type PokedexDescriptions struct {
	Description string              `json:"description"`
	Language    LanguageDescription `json:"language"`
}

type LanguageDescription struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonEntries struct {
	EntryNumber int            `json:"entry_number"`
	Pokemon     PokemonSpecies `json:"pokemon_species"`
}

type PokemonSpecies struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Pokedexes : List the Pokemon available from a specified Trainer
func Pokedexes(p string) {
	var url string
	url = "http://pokeapi.co/api/v2/pokedex/" + p

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}
	// 	Print the entire RepsonseData
	// fmt.Println(string(responseData))

	var responseObject Pokedex
	json.Unmarshal(responseData, &responseObject)

	// List the Pokemon Belonging to the trainer
	for i := 0; i < len(responseObject.Pokemon); i++ {
		fmt.Println(responseObject.Pokemon[i].Pokemon.Name)
	}
}
