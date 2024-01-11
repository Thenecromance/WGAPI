package structure

type Gun struct {
	AimTime float64 /*float*/ `json:"aim_time" xml:"aim_time" ` // Aiming time (s)

	Dispersion float64 /*float*/ `json:"dispersion" xml:"dispersion" ` // Dispersion at 100 m (m)

	FireRate float64 /*float*/ `json:"fire_rate" xml:"fire_rate" ` // Rate of fire (rounds/min)

	MaxAmmo int32 /*numeric*/ `json:"max_ammo" xml:"max_ammo" ` // Number of shells

	MoveDownArc int32 /*numeric*/ `json:"move_down_arc" xml:"move_down_arc" ` // Depression angle (deg)

	MoveUpArc int32 /*numeric*/ `json:"move_up_arc" xml:"move_up_arc" ` // Elevation angle (deg)

	ReloadTime float64 /*float*/ `json:"reload_time" xml:"reload_time" ` // Reload time (s)

	TraverseSpeed int32 /*numeric*/ `json:"traverse_speed" xml:"traverse_speed" ` // Traverse speed (deg/s)

	Ammo []*Ammo `json:"ammo" xml:"ammo" ` // Gun shells characteristics
}
