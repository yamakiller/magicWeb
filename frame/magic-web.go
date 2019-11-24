package frame

import (
	"github.com/yamakiller/magicLibs/args"
	"github.com/yamakiller/magicLibs/logger"
)

//IMount desc
//@interface IMount desc: web mount interface
type IMount interface {
	Start() error
	FillRouter() error
	Shutdown()
}

//MagicWeb desc
//@method MagicWeb desc: web system main frame
type MagicWeb struct {
	_mount IMount
}

//Mount desc
//@method Mount desc: mount web system
func (slf *MagicWeb) Mount() {
	var err error
	args.Instance().Parse()
	if err = slf._mount.Start(); err != nil {
		logger.Error(0, "%+v", err)
		goto exit
	}

exit:
	slf._mount.Shutdown()
}
