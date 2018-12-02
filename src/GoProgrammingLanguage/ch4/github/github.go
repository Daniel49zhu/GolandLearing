package github

import "time"

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number   int
	HTMURL   string `json:"html_url"`
	Title    string
	State    string
	User     *user
	CreateAt time.Time `json:"createed_at"`
	Body     string
}

type user struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
