package model

import "time"

type Post struct {
	Tags      []Tags    `json:"tags"`
	Body      string    `json:"body"`
	Title     string    `json:"title"`
	Media     string    `json:"media"`
	About     string    `json:"about"`
	Author    User      `json:"author"`
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	Name    string   `json:"name"`
	Email   string   `json:"email"`
	Image   string   `json:"image"`
	About   string   `json:"about"`
	Posts   []Post   `json:"posts"`
	Socials []string `json:"socials"`
	CreatedAt time.Time `json:"created_at"`
}

type Tags struct {
	Art           bool `json:"art"`
	Tech          bool `json:"tech"`
	Games         bool `json:"games"`
	Books         bool `json:"books"`
	Movies        bool `json:"movies"`
	Travel        bool `json:"travel"`
	Comedy        bool `json:"comedy"`
	Health        bool `json:"health"`
	History       bool `json:"history"`
	Finance       bool `json:"finance"`
	Politics      bool `json:"politics"`
	Science       bool `json:"science"`
	Entertainment bool `json:"entertainment"`
}