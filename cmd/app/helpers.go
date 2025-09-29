package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"time"
)

type staredRepo struct {
	Repos []Repo `json:"items"`
}
type Repo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"html_url"`
}
type User struct {
	Username        string
	Name            string
	Bio             string
	Followers       float64
	Following       float64
	Repo_count      float64
	Profile_pic_url string
	Profile_url     string
	Repos           []Repo
}

var PERSONAL_TOKEN = os.Getenv("GIT_PERSONAL_TOKEN")

func fetchUserInfo(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", PERSONAL_TOKEN))

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	return data, nil
}

func serverError(w http.ResponseWriter, err error, errStr string) {
	log.Println(err)
	// http.Error(w, "Internal server error", http.StatusInternalServerError)
	fmt.Fprintf(os.Stdout, "trace: %s\n", string(debug.Stack()))
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/notfound.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		serverError(w, err, "Internal Server Error")
		return
	}
	err = ts.ExecuteTemplate(w, "base", errStr)
	if err != nil {
		serverError(w, err, "Internal Server Error")
		return
	}
}
func recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				r.Header.Set("Connection", "close")
				// serverError(w, fmt.Errorf("%s", err))
				serverError(w, err.(error), "Internal Server Error")
			}
		}()
		next.ServeHTTP(w, r)
	})
}
