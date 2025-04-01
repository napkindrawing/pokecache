//go:generate go tool go-enum -f=$GOFILE --marshal --flag --nocase --mustparse --names
package models

/*
	ENUM(

Normal
Fire
Water
Electric
Grass
Ice
Fighting
Poison
Ground
Flying
Psychic
Bug
Rock
Ghost
Dragon
Dark
Steel
Fairy
)
*/
type PokemonType string
