package frame

import "github.com/gin-gonic/gin"

//Spawn doc
//@type Spawn @Summary Create agic web framework function
type Spawn func() IMagicWeb

//IMagicWeb doc
//@Method IMagicWeb @Summary eb system main frame
type IMagicWeb interface {
	Start() error
	Shutdown()
	Engine() *gin.Engine
}
