package structure

type Mastery struct {
	Distribution map[string]map[string]int32 /*associative array*/ `json:"distribution" xml:"distribution" ` // Values of these percentiles for each piece of equipment

	UpdatedAt int64 /*timestamp*/ `json:"updated_at" xml:"updated_at" ` // Date of data update

}
