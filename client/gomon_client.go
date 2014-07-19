package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	resource := strings.ToLower(readLine(stdin))

	fmt.Print("Resource ID: ")
	id, err := strconv.Atoi(readLine(stdin))
	check(err)

	var thing interface{}
	switch resource {
	case `pokedex`:
		thing, err = gomon.GetPokedex()
		check(err)
	case `pokemon`:
		thing, err = gomon.GetPokemon(id)
		check(err)
	case `type`:
		thing, err = gomon.GetType(id)
		check(err)
	case `move`:
		thing, err = gomon.GetMove(id)
		check(err)
	case `ability`:
		thing, err = gomon.GetAbility(id)
		check(err)
	case `egg`:
		thing, err = gomon.GetEgg(id)
		check(err)
	case `description`:
		thing, err = gomon.GetDescription(id)
		check(err)
	case `sprite`:
		thing, err = gomon.GetSprite(id)
		check(err)
	case `game`:
		thing, err = gomon.GetGame(id)
		check(err)
	}

	fmt.Printf("%+v", thing)
}
