package auth

import "github.com/dgrijalva/jwt-go"

//Claims desc
//@struct Claims desc: jwt data
type Claims struct {
	Userid   string `json:"userid"`
	Username string `json:"username"`
	Password string `json:"password"`
	IsValid  bool   `json:"valid"`
	jwt.StandardClaims
}
