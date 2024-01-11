package structure

type BuildingSlots struct {
	ArenaId string /*string*/ `json:"arena_id" xml:"arena_id" ` // Map ID associated with the current construction site

	BuildingLevel int32 /*numeric*/ `json:"building_level" xml:"building_level" ` // Level of the structure located on the current construction site

	BuildingTitle string /*string*/ `json:"building_title" xml:"building_title" ` // Name of the structure located on the current construction site

	Direction string /*string*/ `json:"direction" xml:"direction" ` // Bridgehead's name

	Position string /*string*/ `json:"position" xml:"position" ` // Position of the construction site

	ReserveTitle string /*string*/ `json:"reserve_title" xml:"reserve_title" ` // Name of the Reserve earned in the structure located on the current construction site

}
