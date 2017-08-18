package main

import (
	"fmt"
)

type Pokemon struct {
	id   int
	name string
}

type Creator interface {
	CreatePokemon() Pokemon
}

func (p Pokemon) CreatePokemon() Pokemon {
	fmt.Println("Call mongo")
	return Pokemon{}
}

func register(c Creator) {
	c.CreatePokemon()
}

func main() {
	register(Pokemon{})
}
