package security

import "github.com/dgrijalva/jwt-go"

type JwtCustomClaims struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	IP          string `json:"ip"`
	IdModerator bool   `json:"moderator"`
	IdAdmin     bool   `json:"admin"`
	jwt.StandardClaims
}
