package structure
type Boosters struct {
Count int32/*numeric*/ `json:"count" xml:"count" ` // Amount of Personal Reserves

ExpirationTime int64/*timestamp*/ `json:"expiration_time" xml:"expiration_time" ` // Expiration time

State string/*string*/ `json:"state" xml:"state" ` // Status of Personal Reserves. Valid values:

}