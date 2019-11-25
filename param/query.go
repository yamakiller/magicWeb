package param

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetQueryInt desc
//@method GetQueryInt desc: Return URL param
//@param (*gin.Context) http context
//@param (string) url param key
//@param (int) default value
func GetQueryInt(c *gin.Context, key string, def int) int {
	if v, ok := c.GetQuery(key); ok {
		if r, e := strconv.Atoi(v); e != nil {
			return r
		}
	}

	return def
}
