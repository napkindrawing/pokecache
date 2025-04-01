package models

type Pokemon struct {
	ID        PokemonID        `json:"id"`
	Name      string           `json:"name"`
	Type      PokemonType      `json:"type"`
	Height    uint32           `json:"height"`
	Weight    uint             `json:"weight"`
	Abilities []PokemonAbility `json:"abilities"`
}
