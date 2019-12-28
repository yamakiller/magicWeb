package session

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RedisStore(name string, size int, addr, pwd, key string) gin.HandlerFunc {

	store, err := sessions.NewRedisStore(size, "tcp", addr, pwd, []byte(key))
	if err != nil {
		panic(err)
	}

	return sessions.Sessions(name, store)
}

func CookieStore(name, key string) gin.HandlerFunc {
	store := sessions.NewCookieStore([]byte(key))
	return sessions.Sessions(name, store)
}
