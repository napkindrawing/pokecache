package models

type PokemonAbility struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Generation  uint8  `json:"generation"`
}
