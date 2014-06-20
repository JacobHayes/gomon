package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/JacobHayes/gomon"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func readLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	check(scanner.Err())
	line := scanner.Text()
	check(scanner.Err())

	return line
}

func main() {
	stdin := bufio.NewScanner(os.Stdin)

	fmt.Println("Resources:")
	fmt.Println("\tPokedex, Pokemon, Type, Move, Ability, Egg, Description, Sprite, Game")
	fmt.Print("\nChoose a resource: ")
	switch strings.ToLower(readLine(stdin)) {
	case `pokedex`:
		resource, err := gomon.GetPokedex(`1`)
		check(err)

		fmt.Println(resource)
	case `pokemon`:
		resource, err := gomon.GetPokemon(`1`)
		check(err)

		fmt.Println(resource)
	case `type`:
		resource, err := gomon.GetType(`1`)
		check(err)

		fmt.Println(resource)
	case `move`:
		resource, err := gomon.GetMove(`1`)
		check(err)

		fmt.Println(resource)
	case `ability`:
		resource, err := gomon.GetAbility(`1`)
		check(err)

		fmt.Println(resource)
	case `egg`:
		resource, err := gomon.GetEgg(`1`)
		check(err)

		fmt.Println(resource)
	case `description`:
		resource, err := gomon.GetDescription(`1`)
		check(err)

		fmt.Println(resource)
	case `sprite`:
		resource, err := gomon.GetSprite(`1`)
		check(err)

		fmt.Println(resource)
	case `game`:
		resource, err := gomon.GetGame(`1`)
		check(err)

		fmt.Println(resource)
	}
}
