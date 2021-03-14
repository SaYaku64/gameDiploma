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

type activeUsers struct {
	sync.RWMutex
	m map[string]user
}

// ActiveUsers - cache that contains all users, that didn't pressed Logout
var ActiveUsers = activeUsers{
	m: make(map[string]user),
}
