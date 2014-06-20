package gomon

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func GetPokedex(id string) (Pokedex, error) {
	pokedex := new(Pokedex)
	if err := getResource(pokedex_const, id, pokedex); err != nil {
		return Pokedex{}, err
	}

	return *pokedex, nil
}

func GetPokemon(id string) (Pokemon, error) {
	pokemon := new(Pokemon)
	if err := getResource(pokemon_const, id, pokemon); err != nil {
		return Pokemon{}, err
	}

	return *pokemon, nil
}

func GetType(id string) (Type, error) {
	type_var := new(Type)
	if err := getResource(type_const, id, type_var); err != nil {
		return Type{}, err
	}

	return *type_var, nil
}

func GetMove(id string) (Move, error) {
	move := new(Move)
	if err := getResource(move_const, id, move); err != nil {
		return Move{}, err
	}

	return *move, nil
}

func GetAbility(id string) (Ability, error) {
	ability := new(Ability)
	if err := getResource(ability_const, id, ability); err != nil {
		return Ability{}, err
	}

	return *ability, nil
}

func GetEgg(id string) (Egg, error) {
	egg := new(Egg)
	if err := getResource(egg_const, id, egg); err != nil {
		return Egg{}, err
	}

	return *egg, nil
}

func GetDescription(id string) (Description, error) {
	description := new(Description)
	if err := getResource(description_const, id, description); err != nil {
		return Description{}, err
	}

	return *description, nil
}

func GetSprite(id string) (Sprite, error) {
	sprite := new(Sprite)
	if err := getResource(sprite_const, id, sprite); err != nil {
		return Sprite{}, err
	}

	return *sprite, nil
}

func GetGame(id string) (Game, error) {
	game := new(Game)
	if err := getResource(game_const, id, game); err != nil {
		return Game{}, err
	}

	return *game, nil
}

func requestUrl(resource string, id string) string {
	return strings.Join([]string{pokeUrl, resource, id, ``}, `/`)
}

func getJSON(resource string, id string) ([]byte, error) {
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

func getResource(resource string, id string, e entity) error {
	entity_json, err := getJSON(resource, id)
	if err != nil {
		return err
	}

	err = e.unmarshal(entity_json)
	if err != nil {
		return err
	}

	return nil
}
