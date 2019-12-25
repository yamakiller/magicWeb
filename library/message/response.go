package message

//Response doc
//Summary HTTP Response informat json
//Struct Response
//Member (int) error code
//Member (interface{}) message or data
type Response struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
}
