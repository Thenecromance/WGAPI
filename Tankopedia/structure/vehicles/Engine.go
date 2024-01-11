package structure
type Engine struct {
FireChance float64/*float*/ `json:"fire_chance" xml:"fire_chance" ` // Chance of engine fire

Name string/*string*/ `json:"name" xml:"name" ` // Module name

Power int32/*numeric*/ `json:"power" xml:"power" ` // Engine Power (hp)

Tag string/*string*/ `json:"tag" xml:"tag" ` // Module tag

Tier int32/*numeric*/ `json:"tier" xml:"tier" ` // Tier

Weight int32/*numeric*/ `json:"weight" xml:"weight" ` // Weight (kg)

}