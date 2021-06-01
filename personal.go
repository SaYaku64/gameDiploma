package main

import (
	"crypto/sha256"
	"fmt"
)

const (
	quickEmail = "929f903df3aeeae441e711bcbd1d743d42cd1f0364a8d78a3fed131af710"

	mongoConn = "mongodb+srv://admin:159753ghj@nulesocial.aijwv.mongodb.net/nulesocial"

	mongoHost             = "nulesocial.aijwv.mongodb.net"
	mongoUsername         = "admin"
	mongoPassword         = "159753ghj"
	mongoDatabase         = "nulesocial"
	mongoCollUsers        = "users"
	mongoCollInfographics = "infographics"
)

// encode - encodes string with SHA-256 hashing algorithm
func encode(str string) string {
	bytes := sha256.Sum256([]byte(str))
	return fmt.Sprintf("%x", bytes)
}

// checkEncoded - checks if first string is the same as second which is encoded
func checkEncoded(str string, str2 string) bool {
	if encode(str) == str2 {
		return true
	}
	return false
}
