package structure
type Hull struct {
Front int32/*numeric*/ `json:"front" xml:"front" ` // Front (mm)

Rear int32/*numeric*/ `json:"rear" xml:"rear" ` // Rear (mm)

Sides int32/*numeric*/ `json:"sides" xml:"sides" ` // Sides (mm)

}