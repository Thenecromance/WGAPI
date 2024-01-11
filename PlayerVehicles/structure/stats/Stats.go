package structure

type Stats struct {
	AccountId int32/*numeric*/ `json:"account_id" xml:"account_id" ` // Player account ID

	MarkOfMastery int32/*numeric*/ `json:"mark_of_mastery" xml:"mark_of_mastery" ` // Mastery Badges:

	MaxFrags int32/*numeric*/ `json:"max_frags" xml:"max_frags" ` // Maximum destroyed in battle

	MaxXp int32/*numeric*/ `json:"max_xp" xml:"max_xp" ` // Maximum experience per battle

	TankId int32/*numeric*/ `json:"tank_id" xml:"tank_id" ` // Vehicle ID

	Frags map[string]string/*associative array*/ `json:"frags" xml:"frags" ` // Details on vehicles destroyed. This data requires a valid access_token for the specified account.

	InGarage bool/*boolean*/ `json:"in_garage" xml:"in_garage" ` // Availability of vehicle in the Garage. This data requires a valid access_token for the specified account.

	All                *All                `json:"all" xml:"all" `                                 // Overall Statistics
	Clan               *Clan               `json:"clan" xml:"clan" `                               // Clan battles statistics
	Company            *Company            `json:"company" xml:"company" `                         // Tank Company battles statistics
	Epic               *Epic               `json:"epic" xml:"epic" `                               // Statistics in Grand Battles.
	Fallout            *Fallout            `json:"fallout" xml:"fallout" `                         // Fallout statistics.
	Globalmap          *Globalmap          `json:"globalmap" xml:"globalmap" `                     // All battle statistics on the Global Map
	Random             *Random             `json:"random" xml:"random" `                           // Random battles statistics.
	Ranked10x10        *Ranked10x10        `json:"ranked_10x10" xml:"ranked_10x10" `               // Archive of statistics for ranked 10x10 battles.
	RankedBattles      *RankedBattles      `json:"ranked_battles" xml:"ranked_battles" `           // Statistics in Ranked Battles.
	RegularTeam        *RegularTeam        `json:"regular_team" xml:"regular_team" `               // Battle statistics of permanent teams
	StrongholdDefense  *StrongholdDefense  `json:"stronghold_defense" xml:"stronghold_defense" `   // General stats for player's battles in Stronghold defense
	StrongholdSkirmish *StrongholdSkirmish `json:"stronghold_skirmish" xml:"stronghold_skirmish" ` // General stats for player's battles in Skirmishes
	Team               *Team               `json:"team" xml:"team" `                               // Team battles statistics
}
