package structure
type Crewskills struct {
Description string/*string*/ `json:"description" xml:"description" ` // Skill description

IsPerk bool/*boolean*/ `json:"is_perk" xml:"is_perk" ` // Indicates if it is a perk

Name string/*string*/ `json:"name" xml:"name" ` // Skill name

Skill string/*string*/ `json:"skill" xml:"skill" ` // Skill ID

ImageUrl *ImageUrl`json:"image_url" xml:"image_url" ` // URL to skill icon
}