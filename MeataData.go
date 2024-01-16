package wgapi

// response from the WarGaming's meta data
type Meta struct {
	Count int64 `json:"count" binding:"required"`
}

// all the response will be like this
type ResponseBase struct {
	Status string `json:"status"  binding:"required"`
	Meta   Meta   `json:"meta"  binding:"required"`
}
