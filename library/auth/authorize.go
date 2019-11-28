package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

//Enter desc
//@method Enter desc: enter jwt system
//@param (string) token secret
//@param (string) user id
//@param (string) user name
//@param (string) user password
//@param (int)    token expire time unit/minute
//@return (string) token
//@return (error)
func Enter(secret, id, name, pwd string, expire int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(expire) * time.Minute)

	claims := Claims{
		id,
		name,
		pwd,
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

//Verify desc
//@method Verify desc: verify token and returns claims
//@param (string) jwt secret
//@param (string) jwt token
//@return (*Claims)
//@return (error)
func Verify(secret, token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
