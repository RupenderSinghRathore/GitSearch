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
		serverError(w, err, "Internal Server Error")
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	// err = ts.Execute(w, nil)
	if err != nil {
		serverError(w, err, "Internal Server Error")
		return
	}
}
func serveUser(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
	infoUrl := fmt.Sprintf("https://api.github.com/users/%s", username)
	userData, err := fetchUserInfo(infoUrl)
	if err != nil {
		serverError(w, err, "Internal Server Error")
		return
	}
	// repoUrl := fmt.Sprintf("https://api.github.com/users/%s/repos?per_page=5", username)
	repoUrl := fmt.Sprintf(
		"https://api.github.com/search/repositories?q=user:%s&sort=stars&order=desc&per_page=5",
		username,
	)
	repoData, err := fetchUserInfo(repoUrl)
	if err != nil {
		serverError(w, err, "Internal Server Error")
		return
	}
	userInfo := make(map[string]any)
	err = json.Unmarshal(userData, &userInfo)
	if err != nil {
		serverError(w, err, "Internal Server Error")
		return
	}
	if status, ok := userInfo["status"].(string); ok {
		code, err := strconv.Atoi(status)
		if err != nil || code >= 400 {
			serverError(w, err, fmt.Sprintf("\"%s\" user not found", username))
			return
		}
	}

	repoInfo := staredRepo{}
	err = json.Unmarshal(repoData, &repoInfo)
	if err != nil {
		serverError(w, err, "Internal Server Error")
		return
	}
	user := User{
		Username:    userInfo["login"].(string),
		Profile_url: userInfo["html_url"].(string),
		Repo_count:  userInfo["public_repos"].(float64),
		Followers:   userInfo["followers"].(float64),
		Following:   userInfo["following"].(float64),
		Repos:       repoInfo.Repos,
	}
	if userInfo["name"] != nil {
		user.Name = userInfo["name"].(string)
	}
	if userInfo["bio"] != nil {
		user.Bio = userInfo["bio"].(string)
	}
	if userInfo["avatar_url"] != nil {
		user.Profile_pic_url = userInfo["avatar_url"].(string)
	}
	// println("bio:", user.Bio)

	files := []string{
		"./ui/html/pages/user.tmpl",
		"./ui/html/base.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		serverError(w, err, "Internal Server Error")
		return
	}
	err = ts.ExecuteTemplate(w, "base", user)
	if err != nil {
		serverError(w, err, "Internal Server Error")
		return
	}
}
