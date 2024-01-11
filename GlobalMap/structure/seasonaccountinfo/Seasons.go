package structure

type Seasons struct {
	AccountId int32/*numeric*/ `json:"account_id" xml:"account_id" ` // Account ID

	AwardLevel string/*string*/ `json:"award_level" xml:"award_level" ` // Award level

	Battles int32/*numeric*/ `json:"battles" xml:"battles" ` // Battles fought for current clan

	BattlesToAward int32/*numeric*/ `json:"battles_to_award" xml:"battles_to_award" ` // Battles to fight in a current clan to get clan award for the season

	ClanId int32/*numeric*/ `json:"clan_id" xml:"clan_id" ` // Clan ID

	ClanRank int32/*numeric*/ `json:"clan_rank" xml:"clan_rank" ` // Clan rating

	SeasonId string/*string*/ `json:"season_id" xml:"season_id" ` // Season ID

	UpdatedAt int64/*timestamp*/ `json:"updated_at" xml:"updated_at" ` // Date when account data were updated

	VehicleLevel int32/*numeric*/ `json:"vehicle_level" xml:"vehicle_level" ` // Vehicle Tier

}
