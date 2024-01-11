package structure
type Missions struct {
Description string/*string*/ `json:"description" xml:"description" ` // Mission description

Hint string/*string*/ `json:"hint" xml:"hint" ` // Hints

MaxLevel int32/*numeric*/ `json:"max_level" xml:"max_level" ` // Maximum vehicle Tier

MinLevel int32/*numeric*/ `json:"min_level" xml:"min_level" ` // Minimum vehicle Tier

MissionId int32/*numeric*/ `json:"mission_id" xml:"mission_id" ` // Mission ID

Name string/*string*/ `json:"name" xml:"name" ` // Mission name

SetId int32/*numeric*/ `json:"set_id" xml:"set_id" ` // Mission branch ID

Tags []string/*list of strings*/ `json:"tags" xml:"tags" ` // Mission tags

Rewards *Rewards`json:"rewards" xml:"rewards" ` // Rewards grouped by mission conditions
}