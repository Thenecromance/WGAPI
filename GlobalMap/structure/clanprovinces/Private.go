package structure

type Private struct {
	HqConnected bool/*boolean*/ `json:"hq_connected" xml:"hq_connected" ` // Indicates availability of connection to the Headquarters

	IsRevenueLimitExceeded bool/*boolean*/ `json:"is_revenue_limit_exceeded" xml:"is_revenue_limit_exceeded" ` // Province income limit. Valid values:

}
