package structure
type Reward struct {
Slots int32/*numeric*/ `json:"slots" xml:"slots" ` // Slots

Tanks []int32/*list of integers*/ `json:"tanks" xml:"tanks" ` // Vehicle IDs

}