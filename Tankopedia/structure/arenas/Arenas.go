package structure
type Arenas struct {
ArenaId string/*string*/ `json:"arena_id" xml:"arena_id" ` // Map ID

CamouflageType string/*string*/ `json:"camouflage_type" xml:"camouflage_type" ` // Map type

Description string/*string*/ `json:"description" xml:"description" ` // Short map description

NameI18n string/*string*/ `json:"name_i18n" xml:"name_i18n" ` // Localized map name

}