package structure
type Armor struct {
Hull *Hull`json:"hull" xml:"hull" ` // Hull armor
Turret *Turret`json:"turret" xml:"turret" ` // Turret armor
}