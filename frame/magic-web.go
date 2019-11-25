package frame

//Spawn desc
//@type Spawn desc: Create magic web framework function
type Spawn func() IMagicWeb

//IMagicWeb desc
//@method IMagicWeb desc: web system main frame
type IMagicWeb interface {
	Start() error
	Shutdown()
}
