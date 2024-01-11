package structure
type Ammo struct {
Damage []int32/*list of integers*/ `json:"damage" xml:"damage" ` // Damage (hp), a list of values: min, avg, max

Penetration []int32/*list of integers*/ `json:"penetration" xml:"penetration" ` // Penetration (mm), a list of values: min, avg, max

Type string/*string*/ `json:"type" xml:"type" ` // Shell type

Stun *Stun`json:"stun" xml:"stun" ` // Stun characteristics
}