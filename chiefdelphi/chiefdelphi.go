package chiefdelphi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

const baseURL = "https://www.chiefdelphi.com"

type User struct {
	ID int
	Username string
	Team int
}

type Post struct {
	ID int
	UserID int
	TopicID int
	Timestamp time.Time
	Body string
}

type Topic struct {
	ID int
	Title string
}

// Make a request to the Chief Delphi Discourse API (see https://docs.discourse.org).
// This helper function will automatically deserialize the response to the appropriate
// struct.
func makeRequest(endpoint string, v interface{}) error {
	response, err := http.Get(baseURL + endpoint)

	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return errors.New("non-200 HTTP status code")
	}

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}

	json.Unmarshal(data, v)

	return nil
}

func sleep() {
	// For some reason, I get corrupt data if I don't sleep.
	time.Sleep(500 * time.Millisecond)
}