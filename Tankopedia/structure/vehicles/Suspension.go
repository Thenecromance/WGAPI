package structure
type Suspension struct {
LoadLimit int32/*numeric*/ `json:"load_limit" xml:"load_limit" ` // Load limit (kg)

Name string/*string*/ `json:"name" xml:"name" ` // Module name

SteeringLockAngle int32/*numeric*/ `json:"steering_lock_angle" xml:"steering_lock_angle" ` // Maximum wheel turning angle (for wheeled vehicles)

Tag string/*string*/ `json:"tag" xml:"tag" ` // Module tag

Tier int32/*numeric*/ `json:"tier" xml:"tier" ` // Tier

TraverseSpeed int32/*numeric*/ `json:"traverse_speed" xml:"traverse_speed" ` // Traverse speed (deg/s)

Weight int32/*numeric*/ `json:"weight" xml:"weight" ` // Weight (kg)

}