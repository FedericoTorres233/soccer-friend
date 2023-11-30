package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/federicotorres233/soccer-friend/types"

	tele "gopkg.in/telebot.v3"
)

func get_last_matches(b *tele.Bot, API_KEY string) {

	b.Handle("/last", func(c tele.Context) error {

		if c.Args() == nil {
			return c.Send("Get the last 5 matches of a team\nUsage: /last [team]")
		}
		team := strings.ToLower(strings.Join(c.Args(), " "))

		var url1 string = fmt.Sprintf("https://apiv2.allsportsapi.com/football/?met=Teams&teamName=%v&APIkey=%v", team, API_KEY)
		resp, err := http.Get(url1)
		if err != nil {
			log.Println(err)
			return c.Send("An error has occurred")
		}
		defer resp.Body.Close()
		body, err1 := ioutil.ReadAll(resp.Body)
		if err1 != nil {
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
		var url string = fmt.Sprintf("https://apiv2.allsportsapi.com/football/?met=H2H&APIkey=%v&firstTeamId=%v&secondTeamId=%v", API_KEY, team_id, team_id+1)
		resp1, err3 := http.Get(url)
		if err3 != nil {
			log.Println(err3)
			return c.Send("An error has occurred")
		}

		defer resp1.Body.Close()
		body1, err4 := ioutil.ReadAll(resp1.Body)
		if err4 != nil {
			log.Println(err4)
			return c.Send("An error has occurred")
		}

		var data1 types.ApiResponseH2H
		err5 := json.Unmarshal(body1, &data1)
		if err5 != nil {
			log.Println(err5)
			return c.Send("An error has occurred")
		}

		last_5_matches := "These are the last 5 matches:\n\n"
		for k, v := range data1.Result.FirstTeamResults {
			if k > 4 {
				break
			}
			if strings.ToLower(v.Event_home_team) == team {
				// Team played at home
				if v.Event_final_result[0] > v.Event_final_result[4] {
					// Won
					last_5_matches += fmt.Sprintf("Won %v versus %v at home \n", v.Event_final_result, v.Event_away_team)
				} else if v.Event_final_result[0] < v.Event_final_result[4] {
					// Lost
					last_5_matches += fmt.Sprintf("Lost %v versus %v at home \n", v.Event_final_result, v.Event_away_team)
				} else {
					// draw
					last_5_matches += fmt.Sprintf("Draw %v versus %v at home \n", v.Event_final_result, v.Event_away_team)
				}
			} else {
				// Team played away
				if v.Event_final_result[0] > v.Event_final_result[4] {
					// Lost
					last_5_matches += fmt.Sprintf("Lost %v versus %v at home \n", v.Event_final_result, v.Event_home_team)
				} else if v.Event_final_result[0] < v.Event_final_result[4] {
					// Won
					last_5_matches += fmt.Sprintf("Won %v versus %v at home \n", v.Event_final_result, v.Event_home_team)
				} else {
					// draw
					last_5_matches += fmt.Sprintf("Draw %v versus %v at home \n", v.Event_final_result, v.Event_home_team)
				}
			}
		}
		return c.Send(last_5_matches)
	})
}
