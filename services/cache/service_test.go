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
	assert.Equal(t, "Gerblik", got.Name, "(retrieved by name)")
}

func TestDelete(t *testing.T) {
	t.Parallel()

	svc := cache.NewService(10)

	tony := models.Pokemon{
		ID:        333,
		Name:      "Tony Danza",
		Type:      models.PokemonTypeNormal,
		Height:    60,
		Weight:    202,
		Abilities: []models.PokemonAbility{},
	}

	errAdd := svc.Add(tony)
	require.NoError(t, errAdd)

	got := svc.GetByID(333)
	assert.NotNil(t, got, "added pokemon is not nil")
	assert.Equal(t, "Tony Danza", got.Name, "(retrieved by name)")

	errDel := svc.Delete(333)
	require.NoError(t, errDel)

	gotDel := svc.GetByID(333)
	assert.Nil(t, gotDel, "deleted pokemon is nil")
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

	gotAliceByIDBefore := svc.GetByID(1)
	assert.NotNil(t, gotAliceByIDBefore, "earliest pokemon (retrieved by id)")
	assert.Equal(t, "Alice", gotAliceByIDBefore.Name, "earliest pokemon name (retrieved by id)")

	gotAliceByNameBefore := svc.GetByName("Alice")
	assert.NotNil(t, gotAliceByNameBefore, "earliest pokemon (retrieved by name)")
	assert.Equal(t, "Alice", gotAliceByNameBefore.Name, "earliest pokemon name (retrieved by name)")

	errAddDumbo := svc.Add(dumbo)
	require.NoError(t, errAddDumbo)

	assert.Equal(t, 3, svc.Count())

	gotAliceByIDAfter := svc.GetByID(1)
	assert.Nil(t, gotAliceByIDAfter, "earliest pokemon was purged due to capacity (retrieved by id)")

	gotAliceByNameAfter := svc.GetByName("Alice")
	assert.Nil(t, gotAliceByNameAfter, "earliest pokemon was purged due to capacity (retrieved by name)")
}
