package main

import (
	"sync"
)

type token struct {
	Name    string `json:"tokenname"`
	Endless bool   `json:"tokenendless"`
}

type user struct {
	Email       string `json:"email"`
	Username    string `json:"username"`
	Password    string `json:"-"`
	AdminStatus bool   `json:"adminstatus"`
	AdminToken  string `json:"admintoken"`
	Token       token  `json:"token"`
}

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Picture string `json:"picture"`
	Price   string `json:"price"`
}

type comment struct {
	ComTime    string `json:"comtime"`
	ComContent string `json:"comcontent"`
	ComName    string `json:"comname"`
}

type activeUsers struct {
	sync.RWMutex
	m map[string]user
}

// ActiveUsers - cache that contains all users, that didn't pressed Logout
var ActiveUsers = activeUsers{
	m: make(map[string]user),
}

var Busket = map[string]int{}

type bskt struct {
	Title  string `json:"title"`
	Amount int    `json:"amount"`
	Price  int    `json:"price"`
}

var BusketSlice []bskt
