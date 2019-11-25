package test

import (
	"testing"

	"github.com/yamakiller/magicLibs/logger"

	"github.com/yamakiller/magicWeb/boot"
	"github.com/yamakiller/magicWeb/frame"
)

//WebFrameTest test
type WebFrameTest struct {
	frame.DefaultWeb
}

func (slf *WebFrameTest) Boot() error {
	logger.Info(0, "1.Connect DB")
	logger.Info(0, "2.Fill Router")
	return nil
}

func TestFrame(t *testing.T) {
	boot.Lanuch(func() frame.IMagicWeb {
		handle := &WebFrameTest{}
		handle.WithStart(handle.Boot)
		return handle
	})
}
