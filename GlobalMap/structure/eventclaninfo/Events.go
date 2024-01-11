package structure

type Events struct {
	BattleFamePoints int32/*numeric*/ `json:"battle_fame_points" xml:"battle_fame_points" ` // Battle Fame Points

	Battles int32/*numeric*/ `json:"battles" xml:"battles" ` // Battles fought

	EventId string/*string*/ `json:"event_id" xml:"event_id" ` // Event ID

	FamePoints int32/*numeric*/ `json:"fame_points" xml:"fame_points" ` // Total Fame Points

	FamePointsSinceTurn int32/*numeric*/ `json:"fame_points_since_turn" xml:"fame_points_since_turn" ` // Change of Fame Points since last turn calculation

	FrontId string/*string*/ `json:"front_id" xml:"front_id" ` // Front ID

	Rank int32/*numeric*/ `json:"rank" xml:"rank" ` // Current rating

	RankDelta int32/*numeric*/ `json:"rank_delta" xml:"rank_delta" ` // Rating changes during previous turn

	TaskFamePoints int32/*numeric*/ `json:"task_fame_points" xml:"task_fame_points" ` // Fame Points for completing a clan task

	Url string/*string*/ `json:"url" xml:"url" ` // Link to Front

	Wins int32/*numeric*/ `json:"wins" xml:"wins" ` // Victories

}
