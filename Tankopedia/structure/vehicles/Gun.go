package structure
type Gun struct {
AimTime float64/*float*/ `json:"aim_time" xml:"aim_time" ` // Aiming time (s)

Caliber int32/*numeric*/ `json:"caliber" xml:"caliber" ` // Caliber (mm)

Dispersion float64/*float*/ `json:"dispersion" xml:"dispersion" ` // Dispersion at 100 m (m)

FireRate float64/*float*/ `json:"fire_rate" xml:"fire_rate" ` // Rate of fire (rounds/min)

MoveDownArc int32/*numeric*/ `json:"move_down_arc" xml:"move_down_arc" ` // Depression angle (deg)

MoveUpArc int32/*numeric*/ `json:"move_up_arc" xml:"move_up_arc" ` // Elevation angle (deg)

Name string/*string*/ `json:"name" xml:"name" ` // Module name

ReloadTime float64/*float*/ `json:"reload_time" xml:"reload_time" ` // Reload time (s)

Tag string/*string*/ `json:"tag" xml:"tag" ` // Module tag

Tier int32/*numeric*/ `json:"tier" xml:"tier" ` // Tier

TraverseSpeed int32/*numeric*/ `json:"traverse_speed" xml:"traverse_speed" ` // Traverse speed (deg/s)

Weight int32/*numeric*/ `json:"weight" xml:"weight" ` // Weight (kg)

}