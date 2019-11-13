package parts

import (
	"io/ioutil"
	"os"
	"runtime"

	"github.com/robertkrimen/otto"
)

//NewJSStack desc
//@struct NewJSStack desc: create JS script virtual machine
func NewJSStack() *JSStack {
	return &JSStack{otto.New()}
}

//JSStack desc
//@struct JSStack desc: javascript virtual machine
type JSStack struct {
	state *otto.Otto
}

//SetInt desc
//@method SetInt desc: Set the Int variable to the JS script
//@param (string) name
//@param (int)    value
func (slf *JSStack) SetInt(name string, val int) {
	slf.state.Set(name, val)
}

//SetFloat desc
//@method SetFloat desc: Set the Float 32 variable to the JS script
//@param (string)     name
//@param (float32)    value
func (slf *JSStack) SetFloat(name string, val float32) {
	slf.state.Set(name, val)
}

//SetDouble desc
//@method SetDouble desc: Set the Float 64 variable to the JS script
//@param (string)     name
//@param (float64)    value
func (slf *JSStack) SetDouble(name string, val float64) {
	slf.state.Set(name, val)
}

//SetBoolean desc
//@method SetBoolean desc: Set Bool variables for JS scripts
//@param (string)     name
//@param (bool)       value
func (slf *JSStack) SetBoolean(name string, val bool) {
	slf.state.Set(name, val)
}

//SetString desc
//@method SetString desc: Set String variables for JS scripts
//@param (string)     name
//@param (string)       value
func (slf *JSStack) SetString(name string, val string) {
	slf.state.Set(name, val)
}

//SetFunc desc
//@method SetFunc desc: Set the js script to call Go's function
//@param (string)       name
//@param (interface{})  value
func (slf *JSStack) SetFunc(name string, fun interface{}) {
	slf.state.Set(name, fun)
}

//ExecuteScriptFile desc
//@method ExecuteScriptFile desc: Execution script file
//@param   (string) scirpt file path
//@return  (otto.Value) javascript result
//@return  (error) javascript execution error result
func (slf *JSStack) ExecuteScriptFile(filename string) (otto.Value, error) {
	dir, err := os.Getwd()
	if err != nil {
		return otto.Value{}, err
	}

	if runtime.GOOS == "windows" {
		dir += "\\" + filename
	} else {
		dir += "/" + filename
	}

	data, err := ioutil.ReadFile(dir)
	if err != nil {
		return otto.Value{}, err
	}

	return slf.state.Run(data)
}
