package types

type Body_last struct {
	Get      string        `json:"get"`
	Results  int           `json:"results"`
	Response []fixtureInfo `json:"response"`
}

type fixtureInfo struct {
	Fixture fixture `json:"fixture"`
	Teams   teams   `json:"teams"`
	Goals   goals   `json:"goals"`
	Score   score   `json:"score"`
}

type fixture struct {
	Id        int    `json:"id"`
	Referee   string `json:"referee"`
	Timezone  string `json:"timezone"`
	Date      string `json:"date"`
	Timestamp int    `json:"timestamp"`
	Status    status `json:"status"`
}

type status struct {
	Long    string `json:"long"`
	Short   string `json:"short"`
	Elapsed int    `json:"elapsed"`
}

type score struct {
	Halftime  goals `json:"halftime"`
	Fulltime  goals `json:"fulltime"`
	Extratime goals `json:"extratime"`
	Penalty   goals `json:"penalty"`
}

type goals struct {
	Home int `json:"home"`
	Away int `json:"away"`
}

type teams struct {
	Home home `json:"home"`
	Away away `json:"away"`
}

type home struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Logo   string `json:"logo"`
	Winner bool   `json:"winner"`
}

type away struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Logo   string `json:"logo"`
	Winner bool   `json:"winner"`
}
