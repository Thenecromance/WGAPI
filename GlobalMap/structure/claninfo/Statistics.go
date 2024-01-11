package structure

type Statistics struct {
	Battles int32/*numeric*/ `json:"battles" xml:"battles" ` // Battles fought

	Battles10Level int32/*numeric*/ `json:"battles_10_level" xml:"battles_10_level" ` // Battles fought in Absolute division

	Battles6Level int32/*numeric*/ `json:"battles_6_level" xml:"battles_6_level" ` // Battles fought in Medium division

	Battles8Level int32/*numeric*/ `json:"battles_8_level" xml:"battles_8_level" ` // Battles fought in Champion division

	Captures int32/*numeric*/ `json:"captures" xml:"captures" ` // Total number by provinces captured by clan

	Losses int32/*numeric*/ `json:"losses" xml:"losses" ` // Defeats

	ProvincesCount int32/*numeric*/ `json:"provinces_count" xml:"provinces_count" ` // Current number of captured provinces

	Wins int32/*numeric*/ `json:"wins" xml:"wins" ` // Victories

	Wins10Level int32/*numeric*/ `json:"wins_10_level" xml:"wins_10_level" ` // Victories in Absolute division

	Wins6Level int32/*numeric*/ `json:"wins_6_level" xml:"wins_6_level" ` // Victories in Medium division

	Wins8Level int32/*numeric*/ `json:"wins_8_level" xml:"wins_8_level" ` // Victories in Champion division

}
