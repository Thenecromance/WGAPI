package structure
type ModulesTree struct {
IsDefault bool/*boolean*/ `json:"is_default" xml:"is_default" ` // Indicates if the module is basic

ModuleId int32/*numeric*/ `json:"module_id" xml:"module_id" ` // Module ID

Name string/*string*/ `json:"name" xml:"name" ` // Module name

NextModules []int32/*list of integers*/ `json:"next_modules" xml:"next_modules" ` // List of module IDs available after research of the module

NextTanks []int32/*list of integers*/ `json:"next_tanks" xml:"next_tanks" ` // List of vehicle IDs available after research of the module

PriceCredit int32/*numeric*/ `json:"price_credit" xml:"price_credit" ` // Cost in credits

PriceXp int32/*numeric*/ `json:"price_xp" xml:"price_xp" ` // Research cost

Type string/*string*/ `json:"type" xml:"type" ` // Module type

}