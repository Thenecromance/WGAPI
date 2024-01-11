package structure
type Turret struct {
Hp int32/*numeric*/ `json:"hp" xml:"hp" ` // Hit points

Name string/*string*/ `json:"name" xml:"name" ` // Module name

Tag string/*string*/ `json:"tag" xml:"tag" ` // Module tag

Tier int32/*numeric*/ `json:"tier" xml:"tier" ` // Tier

TraverseLeftArc int32/*numeric*/ `json:"traverse_left_arc" xml:"traverse_left_arc" ` // Traverse angle, left (deg)

TraverseRightArc int32/*numeric*/ `json:"traverse_right_arc" xml:"traverse_right_arc" ` // Traverse angle, right (deg)

TraverseSpeed int32/*numeric*/ `json:"traverse_speed" xml:"traverse_speed" ` // Traverse speed (deg/s)

ViewRange int32/*numeric*/ `json:"view_range" xml:"view_range" ` // View range (m)

Weight int32/*numeric*/ `json:"weight" xml:"weight" ` // Weight (kg)

}