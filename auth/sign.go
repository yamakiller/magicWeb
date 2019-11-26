package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//SignIn desc
//@struct SignIn
type SignIn struct {
	_cookie string
	_verify func(URL, v string) bool
}

//SetCookieKey desc
//@method SetCookieKey desc: Setting sign in cookie name
//@param (string) cookie name
func (slf *SignIn) SetCookieKey(name string) {
	slf._cookie = name
}

//WithVerify desc
//@method WithVerify desc: With verify function
//@param (func(URL, v string) bool) with function
func (slf *SignIn) WithVerify(f func(URL, v string) bool) {
	slf._verify = f
}

/*
//WithIn desc
//@method WithIn desc: With sign in function
//@param (func(usr interface{})) with function
func (slf *SignIn) WithIn(f func(usr interface{})) {
	slf._in = f
}*/

//In desc
//@method In desc: sign in
//@param (*User) sign in userinfo
func (slf *SignIn) In(name string,
	pwd string,
	query func(name string) interface{},
	save func(interface{})) {
}

//Verify desc
//@method Verify desc: Verify sign in auth
//@return (gin.HandlerFunc) middle function
func (slf *SignIn) Verify() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Request.Cookie("signin_id"); err == nil {
			value := cookie.Value
			if !slf._verify(c.Request.URL.String(), value) {
				goto end_fail
			}
			c.Next()
			return
		}

	end_fail:
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
	}
}
