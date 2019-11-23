package frame

import (
	"github.com/yamakiller/magicLibs/logger"
)

//IMount desc
//@interface IMount desc: web mount interface
//@function (LoadArgs) load command args
//@function (LoadEnv)  load config env
type IMount interface {
	LoadArgs() error
	LoadEnv() error
	Start() error
	Loop() error
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
	if err = slf._mount.LoadArgs(); err != nil {
		panic(err)
	}

	if err = slf._mount.LoadEnv(); err != nil {
		panic(err)
	}

	if err = slf._mount.LoadScript(); err != nil {
		panic(err)
	}

	if err = slf.Start(); err != nil {
		logger.Error(0, "%+v", err)
		goto exit
	}

	if err = slf.Loop(); err != nil {
		logger.Error(0, "%+v", err)
		goto exit
	}

exit:
	slf.Shutdown()
}
