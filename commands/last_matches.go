package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/federicotorres233/soccer-friend/types"
	"github.com/federicotorres233/soccer-friend/utils"

	tele "gopkg.in/telebot.v3"
)

func team_result(a byte, b byte, home bool) string {
	if a == b {
		// draw
		return "d"
	}
	if a > b {
		// won if at home
		if home {
			return "w"
		}
		return "l"
	} else {
		// lost if at home
		if home {
			return "l"
		}
		return "w"
	}
}

func get_last_matches(b *tele.Bot, API_KEY string) {

	b.Handle("/last", func(c tele.Context) error {

		if c.Args() == nil {
			return c.Send("Get the last 5 matches of a team\nUsage: /last [team]")
		}
		team := strings.ToLower(strings.Join(c.Args(), " "))

		url1 := fmt.Sprintf("https://apiv2.allsportsapi.com/football/?met=Teams&teamName=%v&APIkey=%v", team, API_KEY)
		body, err := utils.Fetch(url1)
		if err != nil {
			log.Println(err)
			return c.Send("An error has occurred")
		}

		// Process json data
		var data types.ApiResponseTeams
		err2 := json.Unmarshal(body, &data)
		if err2 != nil {
			log.Println(err)
			return c.Send("An error has occurred")
		}
		if len(data.Result) == 0 {
			return c.Send("Cannot recognize that team, check /teams")
		}

		// Get matches from team id
		team_id := data.Result[0].Key
		url2 := fmt.Sprintf("https://apiv2.allsportsapi.com/football/?met=H2H&APIkey=%v&firstTeamId=%v&secondTeamId=%v", API_KEY, team_id, team_id+1)
		body1, err3 := utils.Fetch(url2)
		if err3 != nil {
			log.Println(err)
			return c.Send("An error has occurred")
		}

		var data1 types.ApiResponseH2H
		err4 := json.Unmarshal(body1, &data1)
		team_results := data1.Result.FirstTeamResults
		if err4 != nil {
			log.Println(err4)
			return c.Send("An error has occurred")
		}
		if len(team_results) == 0 {
			return c.Send("Please enter the full team name, check /teams")
		}

		last_5_matches := "These are the last 5 matches:\n\n"
		for k, v := range team_results {
			if k > 4 {
				break
			}
			if strings.ToLower(v.Event_home_team) == team {
				// Team played at home
				switch team_result(v.Event_final_result[0], v.Event_final_result[4], true) {
				case "w":
					last_5_matches += fmt.Sprintf("Won %v versus %v at home \n", v.Event_final_result, v.Event_away_team)
				case "l":
					last_5_matches += fmt.Sprintf("Lost %v versus %v at home \n", v.Event_final_result, v.Event_away_team)
				case "d":
					last_5_matches += fmt.Sprintf("Draw %v versus %v at home \n", v.Event_final_result, v.Event_away_team)
				}
			} else {
				// Team played away
				switch team_result(v.Event_final_result[0], v.Event_final_result[4], false) {
				case "w":
					last_5_matches += fmt.Sprintf("Won %v versus %v away\n", v.Event_final_result, v.Event_home_team)
				case "l":
					last_5_matches += fmt.Sprintf("Lost %v versus %v away \n", v.Event_final_result, v.Event_home_team)
				case "d":
					last_5_matches += fmt.Sprintf("Draw %v versus %v away \n", v.Event_final_result, v.Event_home_team)
				}
			}
		}
		return c.Send(last_5_matches)
	})
}
