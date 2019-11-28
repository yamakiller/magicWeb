package auth

//SignIn desc
//@struct SignIn
/*type SignIn struct {
	_verify func(URL, v string) bool
}

//WithVerify desc
//@method WithVerify desc: With verify function
//@param (func(URL, v string) bool) with function
func (slf *SignIn) WithVerify(f func(URL, v string) bool) {
	slf._verify = f
}*/

/*
//WithIn desc
//@method WithIn desc: With sign in function
//@param (func(usr interface{})) with function
func (slf *SignIn) WithIn(f func(usr interface{})) {
	slf._in = f
}*/

/*//Enter desc
//@method Enter desc: sign in
//@param (string)  claim key/id
//@param (string)  claim name
//@param (int)     enter time
//@param ([]ClaimPerm) Permission array
//@param (int)  expire time util/Minute
func (slf *SignIn) Enter(key, name string, time, perm []ClaimPerm, expire int) (string, error) {
	loc, _ := time.LoadLocation("PRC")
	//expireTime := nowTime.Add(time)

	return "", nil
}

//Verify desc
//@method Verify desc: Verify sign in auth
//@param  (string) token
//@return (gin.HandlerFunc) middle function
func (slf *SignIn) Verify(token string) gin.HandlerFunc {
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
}*/
