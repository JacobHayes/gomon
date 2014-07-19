package gomon

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func GetPokedex() (Pokedex, error) {
	pokedex := new(Pokedex)
	if err := getResource(pokedexConst, 1, pokedex); err != nil {
		return Pokedex{}, err
	}

	return *pokedex, nil
}

func GetPokemon(id int) (Pokemon, error) {
	pokemon := new(Pokemon)
	if err := getResource(pokemonConst, id, pokemon); err != nil {
		return Pokemon{}, err
	}

	return *pokemon, nil
}

func GetType(id int) (Type, error) {
	type_var := new(Type)
	if err := getResource(typeConst, id, type_var); err != nil {
		return Type{}, err
	}

	return *type_var, nil
}

func GetMove(id int) (Move, error) {
	move := new(Move)
	if err := getResource(moveConst, id, move); err != nil {
		return Move{}, err
	}

	return *move, nil
}

func GetAbility(id int) (Ability, error) {
	ability := new(Ability)
	if err := getResource(abilityConst, id, ability); err != nil {
		return Ability{}, err
	}

	return *ability, nil
}

func GetEgg(id int) (Egg, error) {
	egg := new(Egg)
	if err := getResource(eggConst, id, egg); err != nil {
		return Egg{}, err
	}

	return *egg, nil
}

func GetDescription(id int) (Description, error) {
	description := new(Description)
	if err := getResource(descriptionConst, id, description); err != nil {
		return Description{}, err
	}

	return *description, nil
}

func GetSprite(id int) (Sprite, error) {
	sprite := new(Sprite)
	if err := getResource(spriteConst, id, sprite); err != nil {
		return Sprite{}, err
	}

	return *sprite, nil
}

func GetGame(id int) (Game, error) {
	game := new(Game)
	if err := getResource(gameConst, id, game); err != nil {
		return Game{}, err
	}

	return *game, nil
}

func requestUrl(resource string, id int) string {
	return strings.Join([]string{pokeUrlConst, resource, strconv.Itoa(id), ``}, `/`)
}

func getJSON(resource string, id int) ([]byte, error) {
	resp, err := http.Get(requestUrl(resource, id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	raw_json, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return raw_json, nil
}

func getResource(resource string, id int, e interface{}) error {
	entity_json, err := getJSON(resource, id)
	if err != nil {
		return err
	}

	if len(entity_json) != 0 {
		err = json.Unmarshal(entity_json, e)
		if err != nil {
			return err
		}
	}

	return nil
}
