package structure

type Fronts struct {
	AvgClansRating int32/*numeric*/ `json:"avg_clans_rating" xml:"avg_clans_rating" ` // Average clans rating

	AvgMinBet int32/*numeric*/ `json:"avg_min_bet" xml:"avg_min_bet" ` // Average minimum bid

	AvgWonBet int32/*numeric*/ `json:"avg_won_bet" xml:"avg_won_bet" ` // Average winning bid

	BattleTimeLimit int32/*numeric*/ `json:"battle_time_limit" xml:"battle_time_limit" ` // Maximum battle duration in minutes

	DivisionCost int32/*numeric*/ `json:"division_cost" xml:"division_cost" ` // Division cost

	FogOfWar bool/*boolean*/ `json:"fog_of_war" xml:"fog_of_war" ` // Indicates presence of Fog of War

	FrontId string/*string*/ `json:"front_id" xml:"front_id" ` // Front ID

	FrontName string/*string*/ `json:"front_name" xml:"front_name" ` // Front name

	IsActive bool/*boolean*/ `json:"is_active" xml:"is_active" ` // Indicates if a map is active

	IsEvent bool/*boolean*/ `json:"is_event" xml:"is_event" ` // Indicates the map type: regular map or events map

	MaxTanksPerDivision int32/*numeric*/ `json:"max_tanks_per_division" xml:"max_tanks_per_division" ` // Maximum number of vehicles in division

	MaxVehicleLevel int32/*numeric*/ `json:"max_vehicle_level" xml:"max_vehicle_level" ` // Maximum vehicle Tier

	MinTanksPerDivision int32/*numeric*/ `json:"min_tanks_per_division" xml:"min_tanks_per_division" ` // Minimum number of vehicles in division

	MinVehicleLevel int32/*numeric*/ `json:"min_vehicle_level" xml:"min_vehicle_level" ` // Minimum vehicle Tier

	ProvincesCount int32/*numeric*/ `json:"provinces_count" xml:"provinces_count" ` // Amount of Provinces

	VehicleFreeze bool/*boolean*/ `json:"vehicle_freeze" xml:"vehicle_freeze" ` // Indicates if vehicles blocking is active

	AvailableExtensions *AvailableExtensions `json:"available_extensions" xml:"available_extensions" ` // Available modules
}
