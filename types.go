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

	// AllowedBuildings - cache of info to show
	AllowedBuildings struct {
		Bid  string
		Name string
		Tid  string
	}

	// Building - Building object
	Building struct {
		BID   string
		Added bool
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
		ID       uint64       `json:"_id" bson:"_id"`
		Login    string       `json:"login" bson:"login"`
		Password string       `json:"password" bson:"password"`
		Email    string       `json:"email" bson:"email"`
		Asked    bool         `json:"asked" bson:"asked"`
		Token    token        `json:"token" bson:"token"`
		Fields   []DBBuilding `json:"fields" bson:"fields"`
	}

	// AllSurveys - struct that contain users
	AllSurveys struct {
		Cache map[uint64]Survey
		sync.RWMutex
	}

	// Survey - Survey object
	Survey struct {
		ID        uint64  `json:"_id" bson:"_id"`
		Question1 string  `json:"question1" bson:"question1"`
		Question2 string  `json:"question2" bson:"question2"`
		Question3 int     `json:"question3" bson:"question3"`
		Question4 int     `json:"question4" bson:"question4"`
		Question5 [3]bool `json:"question5" bson:"question5"`
		Question6 [5]bool `json:"question6" bson:"question6"`
	}

	token struct {
		Token   string `json:"token" bson:"token"`
		Endless bool   `json:"endless" bson:"endless"`
	}

	// AllUsers - struct that contain users
	AllUsers struct {
		Cache map[uint64]User
		sync.RWMutex
	}

	DBBuilding struct {
		TID string
		BID string
	}
)
