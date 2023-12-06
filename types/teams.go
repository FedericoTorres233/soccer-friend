package types

type Body_teams struct {
	Get      string     `json:"get"`
	Results  int        `json:"results"`
	Response []teamInfo `json:"response"`
}

type teamInfo struct {
	Team  team  `json:"team"`
	Venue venue `json:"venue"`
}

type team struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Code    string `json:"code"`
	Country string `json:"country"`
}

type venue struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
