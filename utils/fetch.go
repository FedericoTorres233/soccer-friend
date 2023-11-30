package utils

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	// Fetch data from API
	resp, err := http.Get(url)
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
