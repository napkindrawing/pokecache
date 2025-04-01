package models

type Pokemon struct {
	ID   uint64
	Name string
	Type PokemonType
	// Measured in Millimeters
	Height uint32
	// Measured in Milligrams
	Weight    uint
	Abilities []PokemonAbility
}
