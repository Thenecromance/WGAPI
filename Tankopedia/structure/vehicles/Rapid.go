package structure
type Rapid struct {
SpeedBackward int32/*numeric*/ `json:"speed_backward" xml:"speed_backward" ` // Top reverse speed (km/h)

SpeedForward int32/*numeric*/ `json:"speed_forward" xml:"speed_forward" ` // Top speed (km/h)

SuspensionSteeringLockAngle int32/*numeric*/ `json:"suspension_steering_lock_angle" xml:"suspension_steering_lock_angle" ` // Maximum wheel turning angle

SwitchOffTime float64/*float*/ `json:"switch_off_time" xml:"switch_off_time" ` // Time required to switch to Cruise mode

SwitchOnTime float64/*float*/ `json:"switch_on_time" xml:"switch_on_time" ` // Time required to switch to Rapid mode

}