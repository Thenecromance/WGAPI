package structure

type Seasons struct {
	Battles int32 /*numeric*/ `json:"battles" xml:"battles" ` // Battles fought

	Elo int32 /*numeric*/ `json:"elo" xml:"elo" ` // Elo rating

	Rank int32 /*numeric*/ `json:"rank" xml:"rank" ` // Current rating

	RankDelta int32 /*numeric*/ `json:"rank_delta" xml:"rank_delta" ` // Rating changes during previous turn

	VehicleLevel int32 /*numeric*/ `json:"vehicle_level" xml:"vehicle_level" ` // Vehicle Tier

	VictoryPoints int32 /*numeric*/ `json:"victory_points" xml:"victory_points" ` // Victory Points

	VictoryPointsSinceTurn int32 /*numeric*/ `json:"victory_points_since_turn" xml:"victory_points_since_turn" ` // Change of Victory Points since last turn calculation

	Wins int32 /*numeric*/ `json:"wins" xml:"wins" ` // Victories

}
