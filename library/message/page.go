package message

//PageRequest page request info
type PageRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

//PageResponse page response info
type PageResponse struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
}
