package structure

type Info struct {
	GameVersion string/*string*/ `json:"game_version" xml:"game_version" ` // Game client version

	Languages map[string]string/*associative array*/ `json:"languages" xml:"languages" ` // List of supported languages

	TanksUpdatedAt int64/*timestamp*/ `json:"tanks_updated_at" xml:"tanks_updated_at" ` // Time when vehicles details in Tankopedia were updated

	VehicleCrewRoles map[string]string/*associative array*/ `json:"vehicle_crew_roles" xml:"vehicle_crew_roles" ` // Available crew qualifications

	VehicleNations map[string]string/*associative array*/ `json:"vehicle_nations" xml:"vehicle_nations" ` // Nations available

	VehicleTypes map[string]string/*associative array*/ `json:"vehicle_types" xml:"vehicle_types" ` // Available vehicle types

	AchievementSections map[string]AchievementSections `json:"achievement_sections" xml:"achievement_sections" ` // Award sections
}
