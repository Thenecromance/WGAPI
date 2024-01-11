package structure

type Clanprovinces struct {
	ArenaId string/*string*/ `json:"arena_id" xml:"arena_id" ` // Map ID

	ArenaName string/*string*/ `json:"arena_name" xml:"arena_name" ` // Localized map name

	ClanId int32/*numeric*/ `json:"clan_id" xml:"clan_id" ` // Clan ID

	DailyRevenue int32/*numeric*/ `json:"daily_revenue" xml:"daily_revenue" ` // Daily income

	FrontId string/*string*/ `json:"front_id" xml:"front_id" ` // Front ID

	FrontName string/*string*/ `json:"front_name" xml:"front_name" ` // Front name

	LandingType string/*string*/ `json:"landing_type" xml:"landing_type" ` // Indicates the landing type of a province

	MaxVehicleLevel int32/*numeric*/ `json:"max_vehicle_level" xml:"max_vehicle_level" ` // Maximum vehicle Tier in a Front

	PillageEndAt string/*string*/ `json:"pillage_end_at" xml:"pillage_end_at" ` // Date when province will restore its revenue after ransack

	PrimeTime string/*string*/ `json:"prime_time" xml:"prime_time" ` // Prime Time in UTC

	ProvinceId string/*string*/ `json:"province_id" xml:"province_id" ` // Province ID

	ProvinceName string/*string*/ `json:"province_name" xml:"province_name" ` // Province name

	RevenueLevel int32/*numeric*/ `json:"revenue_level" xml:"revenue_level" ` // Income level from 0 to 11. 0 value means that income was not raised.

	TurnsOwned int32/*numeric*/ `json:"turns_owned" xml:"turns_owned" ` // Province owned (number of turns)

	Private *Private `json:"private" xml:"private" ` // Restricted province information
}
