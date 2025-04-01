package cache_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"napkindrawing.com/pokecache/models"
	"napkindrawing.com/pokecache/services/cache"
)

func TestGetByID(t *testing.T) {
	t.Parallel()

	svc := cache.NewService(10)

	buddy := svc.GetByID(1234)
	assert.Nil(t, buddy, "nonexistent pokemon is nil")
}

func TestAdd(t *testing.T) {
	t.Parallel()

	svc := cache.NewService(10)

	gerblik := models.Pokemon{
		ID:        444,
		Name:      "Gerblik",
		Type:      models.PokemonTypeBug,
		Height:    59,
		Weight:    201,
		Abilities: []models.PokemonAbility{},
	}

	errAdd := svc.Add(gerblik)
	require.NoError(t, errAdd)

	got := svc.GetByID(444)
	assert.NotNil(t, got, "added pokemon is not nil")
}

func TestCapacity(t *testing.T) {
	t.Parallel()

	svc := cache.NewService(3)

	alice := models.Pokemon{
		ID:        1,
		Name:      "Alice",
		Type:      models.PokemonTypeBug,
		Height:    101,
		Weight:    201,
		Abilities: []models.PokemonAbility{},
	}

	bob := models.Pokemon{
		ID:        2,
		Name:      "Bob",
		Type:      models.PokemonTypeDragon,
		Height:    102,
		Weight:    202,
		Abilities: []models.PokemonAbility{},
	}

	carol := models.Pokemon{
		ID:        3,
		Name:      "Carol",
		Type:      models.PokemonTypeElectric,
		Height:    103,
		Weight:    203,
		Abilities: []models.PokemonAbility{},
	}

	dumbo := models.Pokemon{
		ID:        4,
		Name:      "Dumbo",
		Type:      models.PokemonTypeFlying,
		Height:    104,
		Weight:    204,
		Abilities: []models.PokemonAbility{},
	}

	assert.Equal(t, 0, svc.Count())

	errAddAlice := svc.Add(alice)
	require.NoError(t, errAddAlice)

	assert.Equal(t, 1, svc.Count())

	errAddBob := svc.Add(bob)
	require.NoError(t, errAddBob)

	assert.Equal(t, 2, svc.Count())

	errAddCarol := svc.Add(carol)
	require.NoError(t, errAddCarol)

	assert.Equal(t, 3, svc.Count())

	errAddDumbo := svc.Add(dumbo)
	require.NoError(t, errAddDumbo)

	assert.Equal(t, 3, svc.Count())

	gotAlice := svc.GetByID(1)
	assert.Nil(t, gotAlice, "earliest pokemon was purged due to capacity")
}
