package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)
// A Response struct to map the Entire Response
type StarWarsPeople struct {
	Name      string   `json:"name"`
	Height    string   `json:"height"`
	Mass      string   `json:"mass"`
	HairColor string   `json:"hair_color"`
	SkinColor string   `json:"skin_color"`
	EyeColor  string   `json:"eye_color"`
	BirthYear string   `json:"birth_year"`
	Gender    string   `json:"gender"`
	Filsm     []string `json:"films"`
}

func main() {
	response, _ := http.Get("https://swapi.dev/api/people/1")

	responseData, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var LukeSkywalker StarWarsPeople
	json.Unmarshal(responseData, &LukeSkywalker)

	fmt.Println("---- Print Field ----")
	fmt.Println(LukeSkywalker.Name)
	fmt.Println(LukeSkywalker.Height)
	fmt.Println(LukeSkywalker.Mass)
	fmt.Println(LukeSkywalker.HairColor)
	fmt.Println(LukeSkywalker.Filsm[0])
  }
  

