package utils

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func Fetch(url string) ([]byte, error) {
	// Fetch data from API
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
  resp.Header.Add("x-rapidapi-key", os.Getenv("API_KEY"))
  resp.Header.Add("x-rapidapi-host", "v3.football.api-sports.io")

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return body, nil

}
