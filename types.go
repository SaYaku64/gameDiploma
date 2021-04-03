package main

import (
	"sync"
)

type (
	// Decoration - decoration object
	Decoration struct {
		DID   string
		DType string
	}

	// BuildingInfoType - cache of info to show
	BuildingInfoType struct {
		Name        string
		Description string
	}

	// Building - Building object
	Building struct {
		BID   string
		BType string
		BSize int
		Info  []BuildingInfoType
	}

	// Territory - territory object
	Territory struct {
		TID   string
		TSize int // 1-5: 1 - the smallest, 5 - the biggest
		Allow bool
	}

	// User - user object
	User struct {
		ID       uint64 `json:"_id"`
		Login    string `json:"login"`
		Password string `json:"password"`
		Email    string `json:"email"`
		Banned   bool   `json:"banned"`
		token
	}

	token struct {
		Token   string `json:"token"`
		Endless bool   `json:"endless"`
	}

	// AllUsers - struct that contain users
	AllUsers struct {
		Cache map[uint64]User
		sync.RWMutex
	}

	// // ActiveUsers - struct that contain map of active users
	// ActiveUsers struct {
	// 	Cache map[string]User
	// 	sync.RWMutex
	// }

	ginError struct {
		Title   string
		Message string
	}
)
