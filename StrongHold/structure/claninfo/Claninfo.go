package structure

type Claninfo struct {
	ClanId int32 /*numeric*/ `json:"clan_id" xml:"clan_id" ` // Clan ID

	ClanName string /*string*/ `json:"clan_name" xml:"clan_name" ` // Clan name

	ClanTag string /*string*/ `json:"clan_tag" xml:"clan_tag" ` // Clan tag

	CommandCenterArenaId string /*string*/ `json:"command_center_arena_id" xml:"command_center_arena_id" ` // Map ID associated with the Command Center construction site

	StrongholdBuildingsLevel int32 /*numeric*/ `json:"stronghold_buildings_level" xml:"stronghold_buildings_level" ` // Total level of the Stronghold's structures

	StrongholdLevel int32 /*numeric*/ `json:"stronghold_level" xml:"stronghold_level" ` // Stronghold's level

	BattlesForStrongholdsStatistics       *BattlesForStrongholdsStatistics       `json:"battles_for_strongholds_statistics" xml:"battles_for_strongholds_statistics" `               // Statistics for the clan's battles in the Stronghold mode
	BattlesSeriesForStrongholdsStatistics *BattlesSeriesForStrongholdsStatistics `json:"battles_series_for_strongholds_statistics" xml:"battles_series_for_strongholds_statistics" ` // Statistics for skirmishes against the clan's Stronghold
	BuildingSlots                         []*BuildingSlots                       `json:"building_slots" xml:"building_slots" `                                                       // Information about the Stronghold's construction sites
	SkirmishStatistics                    *SkirmishStatistics                    `json:"skirmish_statistics" xml:"skirmish_statistics" `                                             // Clan's skirmish statistics
}
