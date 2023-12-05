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

func process_team(c tele.Context, league string) error {
	log.Println(league)
	var team_key string
	switch strings.ToLower(league) {
	case "premier", "premierleague":
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
	url := fmt.Sprintf("https://apiv2.allsportsapi.com/football/?met=Teams&leagueId=%v&APIkey=%v", team_key)
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
}

// /teams command
func get_teams(b *tele.Bot) {

	/*
	  Set up inline keyboard
	*/
	var (
		// Reference to ReplyMarkup
		menu = &tele.ReplyMarkup{ResizeKeyboard: true}

		// Reply buttons.
		premierBtn    = menu.Text("🏴󠁧󠁢󠁥󠁮󠁧󠁿 Premier League")
		bundesligaBtn = menu.Text("🇩🇪 Bundesliga")
		serieaBtn     = menu.Text("🇮🇹 Serie A")
		laligaBtn     = menu.Text("🇪🇸 La Liga")
		ligue1Btn     = menu.Text("🇫🇷 Ligue 1")
		eredivisieBtn = menu.Text("🇳🇱 Eredivisie")
	)

	league_buttons := []tele.Btn{
		premierBtn,
		bundesligaBtn,
		serieaBtn,
		laligaBtn,
		ligue1Btn,
		eredivisieBtn,
	}

	menu.Reply(menu.Split(3, league_buttons)...)

	/*
	  Handle text based input
	*/
	b.Handle("/teams", func(c tele.Context) error {

		if c.Args() == nil {
			return c.Send("Choose a league: /teams [league]\n- Premier League\n- Bundesliga\n- Serie A\n- La Liga\n- Ligue 1\n- Eredivisie", menu)
		}
		league := strings.Join(c.Args(), "")

		return process_team(c, league)

	})

	/*
	   Handle button based input
	*/
	for _, v := range league_buttons {
		b.Handle(&v, func(c tele.Context) error {
			league := strings.ReplaceAll(utils.RemoveEmojis(c.Message().Text), " ", "")
			return process_team(c, league)
		})
	}

}
