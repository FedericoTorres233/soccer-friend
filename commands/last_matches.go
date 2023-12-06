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

func get_last_matches(b *tele.Bot) {

	b.Handle("/last", func(c tele.Context) error {

		if c.Args() == nil {
			return c.Send("Get the last 5 matches of a team\nUsage: /last [team]")
		}
		team := strings.ToLower(strings.Join(c.Args(), " "))

		url1 := fmt.Sprintf("https://v3.football.api-sports.io/teams?search=%v", strings.ToLower(strings.Join(c.Args(), "+")))
		body, err := utils.Fetch(url1)
		if err != nil {
			log.Println(err)
			return c.Send("An error has occurred")
		}

		// Process json data
		var data types.Body_teams
		err2 := json.Unmarshal(body, &data)
		if err2 != nil {
			log.Println(err)
			return c.Send("An error has occurred")
		}
		if len(data.Response) == 0 {
			return c.Send("Cannot recognize that team, check /teams")
		}

		// Get matches from team id
		team_id := data.Response[0].Team.Id
		url2 := fmt.Sprintf("https://v3.football.api-sports.io/fixtures?team=%v&last=5&season=2023", team_id)
		body1, err3 := utils.Fetch(url2)
		if err3 != nil {
			log.Println(err)
			return c.Send("An error has occurred")
		}

		var data1 types.Body_last
		err4 := json.Unmarshal(body1, &data1)
		team_results := data1.Response
		if err4 != nil {
			log.Println(err4)
			return c.Send("An error has occurred")
		}

		last_5_matches := "These are the last 5 matches:\n\n"
		for k, v := range team_results {
			if k > 4 {
				break
			}
			score := fmt.Sprintf("%v - %v", v.Goals.Home, v.Goals.Away)
			if strings.ToLower(v.Teams.Home.Name) == team {
				// team is home
				status := get_match_result(v.Teams.Home.Winner, v.Teams.Away.Winner)
				last_5_matches += fmt.Sprintf("%v %v versus %v at home 🏠\n\n", status, score, v.Teams.Away.Name)
			} else {
				// team is away
				status := get_match_result(v.Teams.Away.Winner, v.Teams.Home.Winner)
				last_5_matches += fmt.Sprintf("%v %v versus %v away ✈️\n\n", status, score, v.Teams.Home.Name)
			}
		}
		return c.Send(last_5_matches)
	})
}

func get_match_result(a bool, b bool) string {
	// a and b will never be equal
	if a {
		return "Won"
	} else if b {
		return "Lost"
	} else {
		return "Draw"
	}
}
