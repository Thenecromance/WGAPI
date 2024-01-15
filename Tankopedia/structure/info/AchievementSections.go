package structure

type AchievementSections struct {
	Name string/*string*/ `json:"name" xml:"name" ` // Award section name

	Order int32/*numeric*/ `json:"order" xml:"order" ` // Award section order

}
