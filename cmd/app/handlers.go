package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/home.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		serverError(w, err)
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	// err = ts.Execute(w, nil)
	if err != nil {
		serverError(w, err)
		return
	}
}
func serveUser(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
	infoUrl := fmt.Sprintf("https://api.github.com/users/%s", username)
	userData, err := fetchUserInfo(infoUrl)
	if err != nil {
		serverError(w, err)
		return
	}
	repoUrl := fmt.Sprintf("https://api.github.com/users/%s/repos?per_page=5", username)
	repoData, err := fetchUserInfo(repoUrl)
	if err != nil {
		print(err)
		serverError(w, err)
		return
	}
	userInfo := make(map[string]any)
	err = json.Unmarshal(userData, &userInfo)
	if err != nil {
		serverError(w, err)
		return
	}
	if status, ok := userInfo["status"].(string); ok {
		code, err := strconv.Atoi(status)
		if err != nil || code >= 400 {
			serverError(w, err)
			return
		}
	}

	repoInfo := make([]Repo, 0, 5)
	err = json.Unmarshal(repoData, &repoInfo)
	if err != nil {
		serverError(w, err)
		return
	}
	fmt.Printf("info: %v\n", userInfo)
	fmt.Printf("repos: %v\n", repoInfo)
	user := User{
		username:    userInfo["login"].(string),
		name:        userInfo["name"].(string),
		profile_url: userInfo["html_url"].(string),
		repo_count:  userInfo["public_repos"].(float64),
		followers:   userInfo["followers"].(float64),
		following:   userInfo["following"].(float64),
		repos:       repoInfo,
	}
	if userInfo["bio"] != nil {
		user.bio = userInfo["bio"].(string)
	}
	if userInfo["avatar_url"] != nil {
		user.profile_pic_url = userInfo["avatar_url"].(string)
	}
	println("bio:", user.bio)
	w.Write(userData)
}
