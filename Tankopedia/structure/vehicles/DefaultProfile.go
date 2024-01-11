package structure
type DefaultProfile struct {
Hp int32/*numeric*/ `json:"hp" xml:"hp" ` // Hit points

HullHp int32/*numeric*/ `json:"hull_hp" xml:"hull_hp" ` // Hull HP

HullWeight int32/*numeric*/ `json:"hull_weight" xml:"hull_weight" ` // Hull weight (kg)

MaxAmmo int32/*numeric*/ `json:"max_ammo" xml:"max_ammo" ` // Ammunition

MaxWeight int32/*numeric*/ `json:"max_weight" xml:"max_weight" ` // Load limit (kg)

SpeedBackward int32/*numeric*/ `json:"speed_backward" xml:"speed_backward" ` // Top reverse speed (km/h)

SpeedForward int32/*numeric*/ `json:"speed_forward" xml:"speed_forward" ` // Top speed (km/h)

Weight int32/*numeric*/ `json:"weight" xml:"weight" ` // Weight (kg)

Ammo *Ammo`json:"ammo" xml:"ammo" ` // Gun shells characteristics
Armor *Armor`json:"armor" xml:"armor" ` // Armor
Engine *Engine`json:"engine" xml:"engine" ` // Engine characteristics
Gun *Gun`json:"gun" xml:"gun" ` // Gun characteristics
Modules *Modules`json:"modules" xml:"modules" ` // Mounted modules
Radio *Radio`json:"radio" xml:"radio" ` // Radio characteristics
Rapid *Rapid`json:"rapid" xml:"rapid" ` // Vehicle characteristics in Rapid mode (for wheeled vehicles)
Siege *Siege`json:"siege" xml:"siege" ` // Vehicle characteristics in Siege mode
Suspension *Suspension`json:"suspension" xml:"suspension" ` // Suspension characteristics
Turret *Turret`json:"turret" xml:"turret" ` // Turret characteristics
}