package structure
type BonusValues struct {
BattleType string/*string*/ `json:"battle_type" xml:"battle_type" ` // Battle type

Value float64/*float*/ `json:"value" xml:"value" ` // Reserve efficiency for each battle type

}