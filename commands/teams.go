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

func get_teams(b *tele.Bot, API_KEY string) {

	b.Handle("/teams", func(c tele.Context) error {

		// Select league
		var team_key string
		if c.Args() == nil {
			return c.Send("Choose a league: /teams [league]\n- Premier\n- Bundesliga\n- SerieA\n- LaLiga\n- Ligue1\n- Eredivisie")
		}
		league := c.Args()[0]
		switch strings.ToLower(league) {
		case "premier":
			team_key = "152"
		case "bundesliga":
			team_key = "175"
		case "seriea":
			team_key = "207"
		case "laliga":
			team_key = "302"
		case "ligue1":
			team_key = "168"
		case "eredivisie":
			team_key = "244"
		default:
			return c.Send("That league does not exist.\nChoose one like this: /teams [league]\n- Premier\n- Bundesliga\n- SerieA\n- LaLiga\n- Ligue1\n- Eredivisie")
		}

		// Fetch data from API
		url := fmt.Sprintf("https://apiv2.allsportsapi.com/football/?met=Teams&leagueId=%v&APIkey=%v", team_key, API_KEY)
		body, err := utils.Fetch(url)
		if err != nil {
			log.Println(err)
			return c.Send("An error has occurred")
		}

		// Process json data
		var data types.ApiResponseTeams
		err1 := json.Unmarshal(body, &data)
		if err1 != nil {
			log.Println(err)
			return c.Send("An error has occurred")
		}

		// Make a string from the team slice
		var teams string
		for k, v := range data.Result {
			if k == len(data.Result)-1 {
				teams += v.Name
				break
			}
			teams += v.Name + ", "
		}

		return c.Send(fmt.Sprintf("These are the teams for %v: %v.\n\nUse /subscribe [team] to receive updates from a team", league, teams))
	})
}
