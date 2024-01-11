package structure

type Private struct {
	DailyWage int32/*numeric*/ `json:"daily_wage" xml:"daily_wage" ` // Influence points spent per day

	Influence int32/*numeric*/ `json:"influence" xml:"influence" ` // Influence points

}
