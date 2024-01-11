package structure

type AvailableExtensions struct {
	Cost int32 /*numeric*/ `json:"cost" xml:"cost" ` // Cost of modules

	DescriptionStrategic string /*string*/ `json:"description_strategic" xml:"description_strategic" ` // Localized description of strategic effect

	DescriptionTactical string /*string*/ `json:"description_tactical" xml:"description_tactical" ` // Localized description of tactical effect

	ExtensionId string /*string*/ `json:"extension_id" xml:"extension_id" ` // Consumable ID

	Name string /*string*/ `json:"name" xml:"name" ` // Localized consumable name

	Wage int32 /*numeric*/ `json:"wage" xml:"wage" ` // Modules upkeep cost

}
