package gomon

const (
	pokeUrlConst     string = `http://pokeapi.co/api/v1`
	pokedexConst     string = `pokedex`
	pokemonConst     string = `pokemon`
	typeConst        string = `type`
	moveConst        string = `move`
	abilityConst     string = `ability`
	eggConst         string = `egg`
	descriptionConst string = `description`
	spriteConst      string = `sprite`
	gameConst        string = `game`
)

type Pokedex struct {
	Created  string `json:"created"`
	Modified string `json:"modified"`
	Name     string `json:"name"`
	Pokemon  []struct {
		Name        string `json:"name"`
		ResourceUri string `json:"resource_uri"`
	} `json:"pokemon"`
	ResourceUri string `json:"resource_uri"`
}

type Pokemon struct {
	Abilities []struct {
		Name        string `json:"name"`
		ResourceUri string `json:"resource_uri"`
	} `json:"abilities"`
	Attack       float64 `json:"attack"`
	CatchRate    float64 `json:"catch_rate"`
	Created      string  `json:"created"`
	Defense      float64 `json:"defense"`
	Descriptions []struct {
		Name        string `json:"name"`
		ResourceUri string `json:"resource_uri"`
	} `json:"descriptions"`
	EggCycles float64 `json:"egg_cycles"`
	EggGroups []struct {
		Name        string `json:"name"`
		ResourceUri string `json:"resource_uri"`
	} `json:"egg_groups"`
	EvYield    string `json:"ev_yield"`
	Evolutions []struct {
		Level       float64 `json:"level"`
		Method      string  `json:"method"`
		ResourceUri string  `json:"resource_uri"`
		To          string  `json:"to"`
	} `json:"evolutions"`
	Exp             float64 `json:"exp"`
	GrowthRate      string  `json:"growth_rate"`
	Happiness       float64 `json:"happiness"`
	Height          string  `json:"height"`
	Hp              float64 `json:"hp"`
	MaleFemaleRatio string  `json:"male_female_ratio"`
	Modified        string  `json:"modified"`
	Moves           []struct {
		LearnType   string `json:"learn_type"`
		Name        string `json:"name"`
		ResourceUri string `json:"resource_uri"`
	} `json:"moves"`
	Name        string  `json:"name"`
	NationalID  float64 `json:"national_id"`
	PkdxID      float64 `json:"pkdx_id"`
	ResourceUri string  `json:"resource_uri"`
	SpAtk       float64 `json:"sp_atk"`
	SpDef       float64 `json:"sp_def"`
	Species     string  `json:"species"`
	Speed       float64 `json:"speed"`
	Sprites     []struct {
		Name        string `json:"name"`
		ResourceUri string `json:"resource_uri"`
	} `json:"sprites"`
	Total float64 `json:"total"`
	Types []struct {
		Name        string `json:"name"`
		ResourceUri string `json:"resource_uri"`
	} `json:"types"`
	Weight string `json:"weight"`
}

type Type struct {
	Created     string  `json:"created"`
	ID          float64 `json:"id"`
	Ineffective []struct {
		Name        string `json:"name"`
		ResourceUri string `json:"resource_uri"`
	} `json:"ineffective"`
	Modified string `json:"modified"`
	Name     string `json:"name"`
	NoEffect []struct {
		Name        string `json:"name"`
		ResourceUri string `json:"resource_uri"`
	} `json:"no_effect"`
	Resistance []struct {
		Name        string `json:"name"`
		ResourceUri string `json:"resource_uri"`
	} `json:"resistance"`
	ResourceUri    string `json:"resource_uri"`
	SuperEffective []struct {
		Name        string `json:"name"`
		ResourceUri string `json:"resource_uri"`
	} `json:"super_effective"`
	Weakness []struct {
		Name        string `json:"name"`
		ResourceUri string `json:"resource_uri"`
	} `json:"weakness"`
}

type Move struct {
	Accuracy    float64 `json:"accuracy"`
	Category    string  `json:"category"`
	Created     string  `json:"created"`
	Description string  `json:"description"`
	ID          float64 `json:"id"`
	Modified    string  `json:"modified"`
	Name        string  `json:"name"`
	Power       float64 `json:"power"`
	Pp          float64 `json:"pp"`
	ResourceUri string  `json:"resource_uri"`
}

type Ability struct {
	Created     string  `json:"created"`
	Description string  `json:"description"`
	ID          float64 `json:"id"`
	Modified    string  `json:"modified"`
	Name        string  `json:"name"`
	ResourceUri string  `json:"resource_uri"`
}

type Egg struct {
	Created  string  `json:"created"`
	ID       float64 `json:"id"`
	Modified string  `json:"modified"`
	Name     string  `json:"name"`
	Pokemon  []struct {
		Name        string `json:"name"`
		ResourceUri string `json:"resource_uri"`
	} `json:"pokemon"`
	ResourceUri string `json:"resource_uri"`
}

type Description struct {
	Created     string `json:"created"`
	Description string `json:"description"`
	Games       []struct {
		Name        string `json:"name"`
		ResourceUri string `json:"resource_uri"`
	} `json:"games"`
	ID       float64 `json:"id"`
	Modified string  `json:"modified"`
	Name     string  `json:"name"`
	Pokemon  struct {
		Name        string `json:"name"`
		ResourceUri string `json:"resource_uri"`
	} `json:"pokemon"`
	ResourceUri string `json:"resource_uri"`
}

type Sprite struct {
	Created  string  `json:"created"`
	ID       float64 `json:"id"`
	Image    string  `json:"image"`
	Modified string  `json:"modified"`
	Name     string  `json:"name"`
	Pokemon  struct {
		Name        string `json:"name"`
		ResourceUri string `json:"resource_uri"`
	} `json:"pokemon"`
	ResourceUri string `json:"resource_uri"`
}

type Game struct {
	Created     string  `json:"created"`
	Generation  float64 `json:"generation"`
	ID          float64 `json:"id"`
	Modified    string  `json:"modified"`
	Name        string  `json:"name"`
	ReleaseYear float64 `json:"release_year"`
	ResourceUri string  `json:"resource_uri"`
}
