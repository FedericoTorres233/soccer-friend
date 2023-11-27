package commands

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/federicotorres233/telebot/types"

	tele "gopkg.in/telebot.v3"
)

func get_teams(b *tele.Bot) {
	var API_KEY string = os.Getenv("APIKEY")

	b.Handle("/teams", func(c tele.Context) error {

		// Select league
		var team_key string
		league := c.Args()
		if c.Args() == nil {
			return c.Send("Choose a league: /teams [league]\n- Premier\n- Bundesliga\n- SerieA\n- LaLiga\n- Ligue1\n- Eredivisie")
		}
		switch strings.ToLower(league[0]) {
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
			c.Send("Wrong league")
			return c.Send("Choose one like this: /teams [league]\n- Premier\n- Bundesliga\n- SerieA\n- LaLiga\n- Ligue1\n- Eredivisie")
		}

		// Fetch data from API
		var url string = "https://apiv2.allsportsapi.com/football/?met=Teams&leagueId=" + team_key + "&APIkey=" + API_KEY
		resp, err := http.Get(url)
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
		var data types.Result
		err2 := json.Unmarshal(body, &data)
		if err2 != nil {
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

		return c.Send(teams)
	})
}
