package types

type Body struct {
	Get      string     `json:"get"`
	Results  int        `json:"results"`
	Response []TeamInfo `json:"response"`
}

type TeamInfo struct {
	Team  Team  `json:"team"`
	Venue Venue `json:"venue"`
}

type Team struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Code    string `json:"code"`
	Country string `json:"country"`
}

type Venue struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
