package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

//Enter doc
//@Method Enter desc: enter jwt system
//@Param (string) token secret
//@Param (string) user id
//@Param (string) user name
//@Param (string) user password
//@Param (int)    token expire time unit/minute
//@Return (string) token
//@Return (error)
func Enter(secret, id, name, pwd string, expire int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(expire) * time.Minute)

	claims := Claims{
		id,
		name,
		pwd,
		false,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, err
}

//Get doc
//@Method Get desc: token to claims
//@Param (string) jwt secret
//@Param (string) jwt token
//@Return (*Claims)
//@Return (error)
func Get(secret, token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok {
			claims.IsValid = tokenClaims.Valid
			return claims, nil
		}
	}

	return nil, err
}

//Verify doc
//@Method Verify desc: verify token and returns claims
//@Param (string) jwt secret
//@Param (string) jwt token
//@Return (*Claims)
//@Return (error)
func Verify(secret, token string) (*Claims, error) {
	tokenClaims, err := Get(secret, token)
	if tokenClaims != nil && tokenClaims.IsValid {
		return tokenClaims, nil
	}
	return nil, err
}
