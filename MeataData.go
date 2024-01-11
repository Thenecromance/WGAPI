package wgapi

type Meta struct {
	Count int64 `json:"count" binding:"required"`
}

type ResponseBase struct {
	Status string `json:"status"  binding:"required"`
	Meta   Meta   `json:"meta"  binding:"required"`
}
