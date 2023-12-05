package utils

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func Fetch(url string) ([]byte, error) {
	// Fetch data from API
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Add("x-rapidapi-key", os.Getenv("API_KEY"))
	req.Header.Add("x-rapidapi-host", "v3.football.api-sports.io")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return body, nil

}
