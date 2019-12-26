package auth

import "github.com/dgrijalva/jwt-go"

//Claims doc
//@Summary jwt data
type Claims struct {
	ID       string `json:"id"`
	Account  string `json:"account"`
	Password string `json:"password"`
	IsValid  bool   `json:"valid"`
	jwt.StandardClaims
}
