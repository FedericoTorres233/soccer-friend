package types

type ApiResponseTeams struct {
	Result  []team `json:"result"`
	Success int    `json:"success"`
}
type team_ struct {
	Name string `json:"team_name"`
	Key  int    `json:"team_key"`
}

type ApiResponseH2H struct {
	Result  result `json:"result"`
	Success int    `json:"success"`
}
type result struct {
	H2H               []match `json:"H2H"`
	FirstTeamResults  []match `json:"firstTeamResults"`
	SecondTeamResults []match `json:"secondTeamResults"`
}
type match struct {
	Event_key             int
	Event_date            string
	Event_time            string
	Event_home_team       string
	Home_team_key         int
	Event_away_team       string
	Away_team_key         int
	Event_halftime_result string
	Event_final_result    string
	Event_status          string
	Country_name          string
	League_name           string
	League_key            int
	League_round          string
	League_season         string
	Event_live            string
	Event_country_key     int
}
