package structure
type Suspension struct {
LoadLimit int32/*numeric*/ `json:"load_limit" xml:"load_limit" ` // Load limit (kg)

TraverseSpeed int32/*numeric*/ `json:"traverse_speed" xml:"traverse_speed" ` // Traverse speed (deg/s)

}