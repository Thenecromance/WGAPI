package structure
type Turret struct {
ArmorFront int32/*numeric*/ `json:"armor_front" xml:"armor_front" ` // Armor: front (mm)

ArmorRear int32/*numeric*/ `json:"armor_rear" xml:"armor_rear" ` // Armor: rear (mm)

ArmorSides int32/*numeric*/ `json:"armor_sides" xml:"armor_sides" ` // Armor: sides (mm)

Hp int32/*numeric*/ `json:"hp" xml:"hp" ` // Hit points

TraverseSpeed int32/*numeric*/ `json:"traverse_speed" xml:"traverse_speed" ` // Traverse speed (deg/s)

ViewRange int32/*numeric*/ `json:"view_range" xml:"view_range" ` // View range (m)

}