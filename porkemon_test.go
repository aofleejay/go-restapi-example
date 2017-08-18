package main

import "testing"
import "fmt"

type MockPokemon struct{}

func (mp MockPokemon) CreatePokemon() Pokemon {
	fmt.Println("Call from mock")
	return Pokemon{}
}

func TestCreatePokemonSuccess(t *testing.T) {
	register(MockPokemon{})
}
