package structure

type Seasonratingneighbors struct {
	AwardLevel string/*string*/ `json:"award_level" xml:"award_level" ` // Award level

	ClanId int32/*numeric*/ `json:"clan_id" xml:"clan_id" ` // Clan ID

	Color string/*string*/ `json:"color" xml:"color" ` // Clan color

	Name string/*string*/ `json:"name" xml:"name" ` // Clan name

	Rank int32/*numeric*/ `json:"rank" xml:"rank" ` // Current rating

	RankDelta int32/*numeric*/ `json:"rank_delta" xml:"rank_delta" ` // Rating changes during previous turn

	Tag string/*string*/ `json:"tag" xml:"tag" ` // Clan tag

	UpdatedAt int64/*timestamp*/ `json:"updated_at" xml:"updated_at" ` // Date of rating calculation

	VictoryPoints int32/*numeric*/ `json:"victory_points" xml:"victory_points" ` // Victory Points

	VictoryPointsToNextAward int32/*numeric*/ `json:"victory_points_to_next_award" xml:"victory_points_to_next_award" ` // Victory Points to get the next award

}
