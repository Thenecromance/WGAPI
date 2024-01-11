package structure
type Engine struct {
FireChance float64/*float*/ `json:"fire_chance" xml:"fire_chance" ` // Chance of engine fire

Power int32/*numeric*/ `json:"power" xml:"power" ` // Engine Power (hp)

}