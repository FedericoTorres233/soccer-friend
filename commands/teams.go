package commands

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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
			return c.Send("Please choose a league")
		}
		switch league[0] {
		case "Premier":
			team_key = "152"
		case "Bundesliga":
			team_key = "175"
		case "SerieA":
			team_key = "207"
		case "LaLiga":
			team_key = "302"
		default:
			return c.Send("Wrong league")
		}

		// Fetch data from API
		var url string = "https://apiv2.allsportsapi.com/football/?met=Teams&leagueId=" + team_key + "&APIkey=" + API_KEY
		resp, err := http.Get(url)
		if err != nil {
			log.Println(err)
		}

		defer resp.Body.Close()
		body, err1 := ioutil.ReadAll(resp.Body)
		if err1 != nil {
			log.Println(err)
		}

		// Process json data
		var data types.Result
		err2 := json.Unmarshal(body, &data)
		if err2 != nil {
			log.Println(err)
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
