package structure
type Tanks struct {
MarkOfMastery int32/*numeric*/ `json:"mark_of_mastery" xml:"mark_of_mastery" ` // Mastery Badges:

TankId int32/*numeric*/ `json:"tank_id" xml:"tank_id" ` // Vehicle ID

Statistics *Statistics`json:"statistics" xml:"statistics" ` // Vehicle statistics
}