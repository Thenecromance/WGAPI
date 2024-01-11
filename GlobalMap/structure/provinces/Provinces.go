package structure

type Provinces struct {
	ArenaId string/*string*/ `json:"arena_id" xml:"arena_id" ` // Map ID

	ArenaName string/*string*/ `json:"arena_name" xml:"arena_name" ` // Localized map name

	Attackers []int32/*list of integers*/ `json:"attackers" xml:"attackers" ` // List of IDs of attacking clans

	BattlesStartAt string/*string*/ `json:"battles_start_at" xml:"battles_start_at" ` // Battles start time in UTC

	Competitors []int32/*list of integers*/ `json:"competitors" xml:"competitors" ` // List of IDs of participating clans

	CurrentMinBet int32/*numeric*/ `json:"current_min_bet" xml:"current_min_bet" ` // Current minimum bid

	DailyRevenue int32/*numeric*/ `json:"daily_revenue" xml:"daily_revenue" ` // Daily income

	FrontId string/*string*/ `json:"front_id" xml:"front_id" ` // Front ID

	FrontName string/*string*/ `json:"front_name" xml:"front_name" ` // Front name

	IsBordersDisabled bool/*boolean*/ `json:"is_borders_disabled" xml:"is_borders_disabled" ` // Province borders are closed

	LandingType string/*string*/ `json:"landing_type" xml:"landing_type" ` // Landing type: auction, tournament or null

	LastWonBet int32/*numeric*/ `json:"last_won_bet" xml:"last_won_bet" ` // Last winning bid

	MaxBets int32/*numeric*/ `json:"max_bets" xml:"max_bets" ` // Maximum number of bids

	Neighbours []string/*list of strings*/ `json:"neighbours" xml:"neighbours" ` // List of adjacent provinces' IDs

	OwnerClanId int32/*numeric*/ `json:"owner_clan_id" xml:"owner_clan_id" ` // Owning clan ID

	PillageEndAt string/*string*/ `json:"pillage_end_at" xml:"pillage_end_at" ` // Date when province will restore its revenue after ransack

	PrimeTime string/*string*/ `json:"prime_time" xml:"prime_time" ` // Prime Time in UTC

	ProvinceId string/*string*/ `json:"province_id" xml:"province_id" ` // Province ID

	ProvinceName string/*string*/ `json:"province_name" xml:"province_name" ` // Province name

	RevenueLevel int32/*numeric*/ `json:"revenue_level" xml:"revenue_level" ` // Income level from 0 to 11. 0 value means that income was not raised.

	RoundNumber int32/*numeric*/ `json:"round_number" xml:"round_number" ` // Round

	Server string/*string*/ `json:"server" xml:"server" ` // Server ID

	Status string/*string*/ `json:"status" xml:"status" ` // Tournament status: STARTED, FINISHED or null

	Uri string/*string*/ `json:"uri" xml:"uri" ` // Relative link to province

	WorldRedivision bool/*boolean*/ `json:"world_redivision" xml:"world_redivision" ` // Indicates if Repartition of the World is active

	ActiveBattles *ActiveBattles `json:"active_battles" xml:"active_battles" ` // Current battles
}
