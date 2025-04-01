package cache

import (
	"errors"
	"fmt"
	"sync"

	"napkindrawing.com/pokecache/models"
)

var (
	ErrNotFound      = errors.New("pokemon not found")
	ErrDuplicateID   = errors.New("duplicate pokemon id")
	ErrDuplicateName = errors.New("duplicate pokemon name")
	ErrInternal      = errors.New("unexpected internal error")
)

type Service interface {
	GetByID(id models.PokemonID) *models.Pokemon
	GetByName(name string) *models.Pokemon
	Add(pokemon models.Pokemon) error
	Delete(id models.PokemonID) error
	Count() int
}

//nolint:ireturn // We want to return an interface here for clarity
func NewService(maxCapacity int) Service {
	svc := &serviceImpl{
		mut:             sync.Mutex{},
		data:            map[string]models.Pokemon{},
		idNameLookup:    map[models.PokemonID]string{},
		maxCapacity:     maxCapacity,
		nameInsertOrder: []string{},
	}

	return svc
}

type serviceImpl struct {
	mut             sync.Mutex
	data            map[string]models.Pokemon
	idNameLookup    map[models.PokemonID]string
	maxCapacity     int
	nameInsertOrder []string
}

func (svc *serviceImpl) Count() int {
	svc.mut.Lock()
	defer svc.mut.Unlock()

	return len(svc.data)
}

func (svc *serviceImpl) GetByID(id models.PokemonID) *models.Pokemon {
	svc.mut.Lock()
	defer svc.mut.Unlock()

	name, okName := svc.idNameLookup[id]
	if !okName {
		return nil
	}

	pokemon, okPokemon := svc.data[name]
	if !okPokemon {
		return nil
	}

	return &pokemon
}

func (svc *serviceImpl) GetByName(name string) *models.Pokemon {
	svc.mut.Lock()
	defer svc.mut.Unlock()

	pokemon, okPokemon := svc.data[name]
	if !okPokemon {
		return nil
	}

	return &pokemon
}

func (svc *serviceImpl) Delete(id models.PokemonID) error {
	svc.mut.Lock()
	defer svc.mut.Unlock()

	name, okName := svc.idNameLookup[id]
	if !okName {
		return fmt.Errorf("%w with id %d", ErrNotFound, id)
	}

	_, okPokemon := svc.data[name]
	if !okPokemon {
		return fmt.Errorf("%w internal data missing for pokemon '%s' from id %d", ErrInternal, name, id)
	}

	delete(svc.data, name)
	delete(svc.idNameLookup, id)

	return nil
}

func (svc *serviceImpl) Add(pokemon models.Pokemon) error {
	svc.mut.Lock()
	defer svc.mut.Unlock()

	name, okName := svc.idNameLookup[pokemon.ID]
	if okName {
		return fmt.Errorf("%w: id %d already in use for pokemon '%s'", ErrDuplicateID, pokemon.ID, name)
	}

	existing, okExisting := svc.data[pokemon.Name]
	if okExisting {
		return fmt.Errorf("%w: name '%s' already in use by pokemon id %d", ErrDuplicateName, pokemon.Name, existing.ID)
	}

	// Uh-oh, are we at capacity?
	if len(svc.data) == svc.maxCapacity {
		// What is the oldest name?
		oldestName := svc.nameInsertOrder[0]
		svc.nameInsertOrder = svc.nameInsertOrder[1:]

		// Always check wacky edge-cases, don't assume anything!
		// The error message emitted below might save us someday.
		oldestPokemon, okOldest := svc.data[oldestName]
		if !okOldest {
			return fmt.Errorf("%w: oldest inserted pokemon name '%s' not found in internal data", ErrInternal, oldestName)
		}

		delete(svc.data, oldestName)
		delete(svc.idNameLookup, oldestPokemon.ID)
	}

	svc.data[pokemon.Name] = pokemon
	svc.idNameLookup[pokemon.ID] = pokemon.Name
	svc.nameInsertOrder = append(svc.nameInsertOrder, pokemon.Name)

	return nil
}
