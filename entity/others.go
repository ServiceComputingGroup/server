package entity

type Film struct {
	Title         string   `json:"title"`
	Episode_id    int      `json:"episode_id"`
	Opening_crawl string   `json:"opening_crawl"`
	Director      string   `json:"director"`
	Producer      string   `json:"producer"`
	Release_date  string   `json:"release_date"`
	Characters    []string `json:"characters"`
	Planets       []string `json:"planets"`
	Starships     []string `json:"starships"`
	Vehicles      []string `json:"vehicles"`
	Species       []string `json:"species"`
	Created       string   `json:"created"`
	Edited        string   `json:"edited"`
	Url           string   `json:"url"`
}
type People struct {
	Name       string   `json:"name"`
	Height     string   `json:"height"`
	Mass       string   `json:"mass"`
	Hair_color string   `json:"hair_color"`
	Skin_color string   `json:"skin_color"`
	Eye_color  string   `json:"eye_color"`
	Birth_year string   `json:"birth_year"`
	Gender     string   `json:"gender"`
	Homeworld  string   `json:"homeworld"`
	Films      []string `json:"films"`
	Planets    []string `json:"planets"`
	Starships  []string `json:"starships"`
	Vehicles   []string `json:"vehicles"`
	Species    []string `json:"species"`
	Created    string   `json:"created"`
	Edited     string   `json:"edited"`
	Url        string   `json:"url"`
}

type Planet struct {
	Name            string   `json:"name"`
	Rotation_period string   `json:"rotation_period"`
	Orbital_period  string   `json:"orbital_period"`
	Diameter        string   `json:"diameter"`
	Climate         string   `json:"climate"`
	Gravity         string   `json:"gravity"`
	Terrain         string   `json:"terrain"`
	Surface_water   string   `json:"surface_water"`
	Population      string   `json:"population"`
	Residents       []string `json:"residents"`
	Films           []string `json:"films"`
	Created         string   `json:"created"`
	Edited          string   `json:"edited"`
	Url             string   `json:"url"`
}
type Starship struct {
	Name                   string   `json:"name"`
	Model                  string   `json:"model"`
	Manufacturer           string   `json:"manufacturer"`
	Cost_in_credits        string   `json:"cost_in_credits"`
	Length                 string   `json:"length"`
	Max_atmosphering_speed string   `json:"max_atmosphering_speed"`
	Crew                   string   `json:"crew"`
	Passengers             string   `json:"passengers"`
	Cargo_capacity         string   `json:"cargo_capacity"`
	Consumables            string   `json:"consumables"`
	Hyperdrive_rating      string   `json:"hyperdrive_rating"`
	MGLT                   string   `json:"MGLT"`
	Starship_class         string   `json:"starship_class"`
	Pilots                 []string `json:"pilots"`
	Films                  []string `json:"films"`
	Created                string   `json:"created"`
	Edited                 string   `json:"edited"`
	Url                    string   `json:"url"`
}
type Vehicle struct {
	Name                   string   `json:"name"`
	Model                  string   `json:"model"`
	Manufacturer           string   `json:"manufacturer"`
	Cost_in_credits        string   `json:"cost_in_credits"`
	Length                 string   `json:"length"`
	Max_atmosphering_speed string   `json:"max_atmosphering_speed"`
	Crew                   string   `json:"crew"`
	Passengers             string   `json:"passengers"`
	Cargo_capacity         string   `json:"cargo_capacity"`
	Consumables            string   `json:"consumables"`
	Vehicle_class          string   `json:"hyperdrive_rating"`
	Pilots                 []string `json:"pilots"`
	Films                  []string `json:"films"`
	Created                string   `json:"created"`
	Edited                 string   `json:"edited"`
	Url                    string   `json:"url"`
}
type Species struct {
	Name             string   `json:"name"`
	Classification   string   `json:"classification"`
	Designation      string   `json:"designation"`
	Average_height   string   `json:"average_height"`
	Skin_colors      string   `json:"skin_colors"`
	Hair_colors      string   `json:"hair_colors"`
	Eye_colors       string   `json:"eye_colors"`
	Average_lifespan string   `json:"average_lifespan"`
	Homeworld        string   `json:"homeworld"`
	Language         string   `json:"language"`
	People           []string `json:"people"`
	Film             []string `json:"film"`
	Created          string   `json:"created"`
	Edited           string   `json:"edited"`
	Url              string   `json:"url"`
}

type QueryPeople struct {
	Count    string   `json:"count"`
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Result   []People `json:"result"`
}

type QueryPlanet struct {
	Count    string   `json:"count"`
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Result   []Planet `json:"result"`
}

type QuerySpecies struct {
	Count    string    `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Result   []Species `json:"result"`
}

type QueryVehicle struct {
	Count    string    `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Result   []Vehicle `json:"result"`
}

type QueryStarship struct {
	Count    string     `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Result   []Starship `json:"result"`
}

type QueryFilm struct {
	Count    string `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Result   []Film `json:"result"`
}
