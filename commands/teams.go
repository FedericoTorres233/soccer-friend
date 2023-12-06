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
	var league_id string
	switch strings.ToLower(league) {
	case "premier", "premierleague":
		league_id = "39"
	case "bundesliga":
    league_id = "78"
	case "seriea":
    league_id = "135"
	case "laliga":
    league_id = "140"
	case "ligue1":
    league_id = "61"
	case "eredivisie":
    league_id = "88"
	default:
		return c.Send("That league does not exist.\nChoose one like this: /teams [league]\n- Premier\n- Bundesliga\n- SerieA\n- LaLiga\n- Ligue1\n- Eredivisie")
	}

	// Fetch data from API
	url := fmt.Sprintf("https://v3.football.api-sports.io/teams?league=%v&season=2023", league_id)
	body, err := utils.Fetch(url)
	if err != nil {
		log.Println(err)
		return c.Send("An error has occurred")
	}

	// Process json data
	var data types.Body
	err1 := json.Unmarshal(body, &data)
	if err1 != nil {
		log.Println(err)
		return c.Send("An error has occurred")
	}

  log.Println(data)

	// Make a string from the team slice
	var teams string
	for k, v := range data.Response {
		if k == len(data.Response)-1 {
			teams += v.Team.Name
			break
		}
		teams += v.Team.Name + ", "
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
