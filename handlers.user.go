// handlers.user.go

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/quickemailverification/quickemailverification-go"
)

// Generates random token
func generateSessionToken() string {
	// hash, _ := HashString("AcCeSs noT deNIeD")
	// return hash
	return fmt.Sprint(time.Now().UnixNano())
}

func checkEmailValidation(email string) string {
	qev := quickemailverification.CreateClient("929f903df3aeeae441e711bcbd1d743d42cd1f0364a8d78a3fed131af710")
	// Need to use Verify instead Sandbox in production
	response, err := qev.Sandbox(email) // Email address which need to be verified
	if err != nil {
		log.Println(err)
		return "Validation failed"
	}
	return response.Result
}
