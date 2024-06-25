package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Pokemon struct {
	Name      string `json:"name"`
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
		} `json:"ability"`
	} `json:"abilities"`
}

func main() {
	/*
	** Url General => https://pokeapi.co/api/v2/
	** Pokemons => https://pokeapi.co/api/v2/pokemon/
	** Pokemon => https://pokeapi.co/api/v2/pokemon/1/
	 */
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/1/")
	if err != nil {
		panic(err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
			return
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var pokemon Pokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Nombre del Pokémon: %s\n", pokemon.Name)

	fmt.Println("Habilidades:")
	for _, ability := range pokemon.Abilities {
		fmt.Printf("- %s\n", ability.Ability.Name)
	}
}
