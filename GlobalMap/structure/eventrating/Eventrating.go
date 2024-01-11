package structure

type Eventrating struct {
	AwardLevel string /*string*/ `json:"award_level" xml:"award_level" ` // Award level

	BattleFamePoints int32 /*numeric*/ `json:"battle_fame_points" xml:"battle_fame_points" ` // Battle Fame Points

	ClanId int32 /*numeric*/ `json:"clan_id" xml:"clan_id" ` // Clan ID

	Color string /*string*/ `json:"color" xml:"color" ` // Clan color

	FamePointsToImproveAward int32 /*numeric*/ `json:"fame_points_to_improve_award" xml:"fame_points_to_improve_award" ` // Amount of Fame Points needed to improve personal award

	Name string /*string*/ `json:"name" xml:"name" ` // Clan name

	Rank int32 /*numeric*/ `json:"rank" xml:"rank" ` // Current rating

	RankDelta int32 /*numeric*/ `json:"rank_delta" xml:"rank_delta" ` // Rating changes during previous turn

	Tag string /*string*/ `json:"tag" xml:"tag" ` // Clan tag

	TaskFamePoints int32 /*numeric*/ `json:"task_fame_points" xml:"task_fame_points" ` // Fame Points for completing a clan task

	TotalFamePoints int32 /*numeric*/ `json:"total_fame_points" xml:"total_fame_points" ` // Total Fame Points

	UpdatedAt int64 /*timestamp*/ `json:"updated_at" xml:"updated_at" ` // Date of rating calculation

}
