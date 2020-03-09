package trainers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Trainer : Available Trainers JSON Respsonse including count and slice of results
type Trainer struct {
	Count   int              `json:"count"`
	Results []TrainerResults `json:"results"`
}

// TrainerResults : Names of Trainers
type TrainerResults struct {
	Name string `json:"name"`
}

// Trainers : List the available Trainers that have a Pokedex available.
func Trainers() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}
	// 	Print the entire RepsonseData
	//fmt.Println(string(responseData))

	// 	Unmarshal the current responseData
	var responseObject Trainer
	json.Unmarshal(responseData, &responseObject)

	for i := 0; i < len(responseObject.Results); i++ {
		fmt.Println(responseObject.Results[i].Name)
	}
}
