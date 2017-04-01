package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://sdaml.club")
	if err != nil {
		log.Panic(err)
	}

	defer resp.Body.Close() // this is cool
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	log.Println(string(body))
	if err = returnsError(); err != nil { // people do this sometimes
		log.Println(err)
	}

	err = returnsCustomError()
	if err != nil {
		log.Println(err)
	}

}

func returnsError() error {
	return errors.New("Example error")
}

var errCustomError = errors.New("custom error")

func returnsCustomError() error {
	return errCustomError
}
