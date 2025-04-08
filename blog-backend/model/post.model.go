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
	Name      string    `bson:"name,validate" json:"name"`
	Email     string    `bson:"email" json:"email"`
	Image     string    `bson:"image" json:"image"`
	Password  string    `bson:"password" json:"password"`
	About     string    `bson:"about" json:"about"`
	Posts     []Post    `bson:"posts" json:"posts"`
	Socials   []string  `bson:"socials" json:"socials"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
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