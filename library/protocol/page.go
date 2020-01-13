package protocol

//PageRequest page request info
type PageRequest struct {
	Page     int `from:"page" url:"page" json:"page"`
	PageSize int `from:"pageSize" url:"pageSize" json:"pageSize"`
}

//PageResponse page response info
type PageResponse struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
}
