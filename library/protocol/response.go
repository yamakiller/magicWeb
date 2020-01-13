package protocol

//Response doc
//Summary HTTP Response informat json
//Struct Response
//Member (int) error code
//Member (interface{}) message or data
type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}
