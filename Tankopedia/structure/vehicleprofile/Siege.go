package structure
type Siege struct {
AimTime float64/*float*/ `json:"aim_time" xml:"aim_time" ` // Aiming time (s)

Dispersion float64/*float*/ `json:"dispersion" xml:"dispersion" ` // Dispersion at 100 m (m)

MoveDownArc int32/*numeric*/ `json:"move_down_arc" xml:"move_down_arc" ` // Depression angle (deg)

MoveUpArc int32/*numeric*/ `json:"move_up_arc" xml:"move_up_arc" ` // Elevation angle (deg)

ReloadTime float64/*float*/ `json:"reload_time" xml:"reload_time" ` // Reload time (s)

SpeedBackward int32/*numeric*/ `json:"speed_backward" xml:"speed_backward" ` // Top reverse speed (km/h)

SuspensionTraverseSpeed int32/*numeric*/ `json:"suspension_traverse_speed" xml:"suspension_traverse_speed" ` // Standard suspension traverse speed

SwitchOffTime float64/*float*/ `json:"switch_off_time" xml:"switch_off_time" ` // Time needed to switch on the Siege mode

SwitchOnTime float64/*float*/ `json:"switch_on_time" xml:"switch_on_time" ` // Time required to switch to Siege mode

}