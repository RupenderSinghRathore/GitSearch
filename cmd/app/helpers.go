package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Repo struct {
	Name        string `json:"name"`
	Discription string `json:"discription"`
	Url         string `json:"url"`
}
type User struct {
	username        string
	name            string
	bio             string
	followers       float64
	following       float64
	repo_count      float64
	profile_pic_url string
	profile_url     string
	repos           []Repo
}

var PERSONAL_TOKEN = os.Getenv("GIT_PERSONAL_TOKEN")

func fetchUserInfo(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", PERSONAL_TOKEN))

	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	return data, nil
}

func serverError(w http.ResponseWriter, err error) {
	log.Println(err)
	http.Error(w, "Internal server error", http.StatusInternalServerError)
}
