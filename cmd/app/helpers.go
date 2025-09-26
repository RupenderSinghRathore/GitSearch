package main

var repo struct {
	name        string
	discription string
	link        string
}
var user struct {
	username        string
	name            string
	bio             string
	followers       int
	following       int
	repo_count      int
	profile_pic_url string
	profile_url     string
	repos           []any
}
repos := []repo
