package structure

type Events struct {
	AccountId int32 /*numeric*/ `json:"account_id" xml:"account_id" ` // Account ID

	AwardLevel string /*string*/ `json:"award_level" xml:"award_level" ` // Award level

	Battles int32 /*numeric*/ `json:"battles" xml:"battles" ` // Battles fought for current clan

	BattlesToAward int32 /*numeric*/ `json:"battles_to_award" xml:"battles_to_award" ` // Battles to fight in a current clan to get clan award for the season

	ClanId int32 /*numeric*/ `json:"clan_id" xml:"clan_id" ` // Clan ID

	ClanRank int32 /*numeric*/ `json:"clan_rank" xml:"clan_rank" ` // Clan rating

	EventId string /*string*/ `json:"event_id" xml:"event_id" ` // Event ID

	FamePoints int32 /*numeric*/ `json:"fame_points" xml:"fame_points" ` // Total Fame Points

	FamePointsSinceTurn int32 /*numeric*/ `json:"fame_points_since_turn" xml:"fame_points_since_turn" ` // Change of Fame Points since last turn calculation

	FamePointsToImproveAward int32 /*numeric*/ `json:"fame_points_to_improve_award" xml:"fame_points_to_improve_award" ` // Amount of Fame Points needed to improve personal award

	FrontId string /*string*/ `json:"front_id" xml:"front_id" ` // Front ID

	Rank int32 /*numeric*/ `json:"rank" xml:"rank" ` // Current rating

	RankDelta int32 /*numeric*/ `json:"rank_delta" xml:"rank_delta" ` // Rating changes during previous turn

	UpdatedAt int64 /*timestamp*/ `json:"updated_at" xml:"updated_at" ` // Date when account data were updated

	Url string /*string*/ `json:"url" xml:"url" ` // Link to Front

}
