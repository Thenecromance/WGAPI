package structure

type Modules struct {
	Image string /*string*/ `json:"image" xml:"image" ` // Image link

	ModuleId int32 /*numeric*/ `json:"module_id" xml:"module_id" ` // Module ID

	Name string /*string*/ `json:"name" xml:"name" ` // Module name

	Nation string /*string*/ `json:"nation" xml:"nation" ` // Nation

	PriceCredit int32 /*numeric*/ `json:"price_credit" xml:"price_credit" ` // Cost in credits

	Tanks []int32 /*list of integers*/ `json:"tanks" xml:"tanks" ` // Vehicles compatible with module

	Tier int32 /*numeric*/ `json:"tier" xml:"tier" ` // Tier

	Type string /*string*/ `json:"type" xml:"type" ` // Module type

	Weight int32 /*numeric*/ `json:"weight" xml:"weight" ` // Weight (kg)

	DefaultProfile *DefaultProfile `json:"default_profile" xml:"default_profile" ` // Basic technical characteristics of module.
}
