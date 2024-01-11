package structure

type Achievements struct {
	AccountId int32 /*numeric*/ `json:"account_id" xml:"account_id" ` // Player account ID

	Achievements map[string]int32 /*associative array*/ `json:"achievements" xml:"achievements" ` // Achievements earned

	MaxSeries map[string]int32 /*associative array*/ `json:"max_series" xml:"max_series" ` // Maximum values of achievement series

	Series map[string]int32 /*associative array*/ `json:"series" xml:"series" ` // Current values of Achievement Series

	TankId int32 /*numeric*/ `json:"tank_id" xml:"tank_id" ` // Vehicle ID

}
