package gomon

import (
	"testing"
)

const resource string = `pokemon`
const resource_id int = 1
const resource_url string = `http://pokeapi.co/api/v1/pokemon/1/`

func TestRequestUrl(t *testing.T) {
	if u := requestUrl(resource, resource_id); u != resource_url {
		t.Errorf("\nrequestUrl(`%v`, `%v`) returned (%v).\n\tExpected (%v).", resource, resource_id, u, resource_url)
	}
}
