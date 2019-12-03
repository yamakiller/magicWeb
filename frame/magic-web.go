package frame

import "github.com/gin-gonic/gin"

//Spawn desc
//@type Spawn desc: Create agic web framework function
type Spawn func() IMagicWeb

//IMagicWeb desc
//@Method IMagicWeb desc: eb system main frame
type IMagicWeb interface {
	Start() error
	Shutdown()
	Engine() *gin.Engine
}
