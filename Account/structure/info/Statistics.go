package structure
type Statistics struct {
Frags map[string]string/*associative array*/ `json:"frags" xml:"frags" ` // Number and models of vehicles destroyed by a player. Player's private data.

TreesCut int32/*numeric*/ `json:"trees_cut" xml:"trees_cut" ` // Trees knocked down

All *All`json:"all" xml:"all" ` // Total statistics in Random and clan battles without the Global Map 2.0 statistics
Clan *Clan`json:"clan" xml:"clan" ` // Clan battles statistics
Company *Company`json:"company" xml:"company" ` // Tank Company battles statistics
Epic *Epic`json:"epic" xml:"epic" ` // Statistics in Grand Battles.
Fallout *Fallout`json:"fallout" xml:"fallout" ` // Fallout statistics.
GlobalmapAbsolute *GlobalmapAbsolute`json:"globalmap_absolute" xml:"globalmap_absolute" ` // Battle statistics on the Global Map in Absolute division.
GlobalmapChampion *GlobalmapChampion`json:"globalmap_champion" xml:"globalmap_champion" ` // Battle statistics on the Global Map in Champion division.
GlobalmapMiddle *GlobalmapMiddle`json:"globalmap_middle" xml:"globalmap_middle" ` // Battle statistics on the Global Map in Medium division.
Historical *Historical`json:"historical" xml:"historical" ` // Historical battles statistics
Random *Random`json:"random" xml:"random" ` // Random battles statistics.
Ranked10x10 *Ranked10x10`json:"ranked_10x10" xml:"ranked_10x10" ` // Archive of statistics for ranked 10x10 battles.
Ranked15x15 *Ranked15x15`json:"ranked_15x15" xml:"ranked_15x15" ` // Archive of statistics for ranked 15x15 battles.
RankedBattles *RankedBattles`json:"ranked_battles" xml:"ranked_battles" ` // Ranked Battles statistics.
RankedBattlesCurrent *RankedBattlesCurrent`json:"ranked_battles_current" xml:"ranked_battles_current" ` // Current Ranked Battles statistics.
RankedBattlesPrevious *RankedBattlesPrevious`json:"ranked_battles_previous" xml:"ranked_battles_previous" ` // Previous Ranked Battles statistics.
RankedSeason1 *RankedSeason1`json:"ranked_season_1" xml:"ranked_season_1" ` // Statistics of ranked battles for the first season of the year. Updated with every battle this season, data is reseted at the end of the year..
RankedSeason2 *RankedSeason2`json:"ranked_season_2" xml:"ranked_season_2" ` // Statistics of ranked battles for the second season of the year. Updated with every battle this season, data is reseted at the end of the year..
RankedSeason3 *RankedSeason3`json:"ranked_season_3" xml:"ranked_season_3" ` // Statistics of ranked battles for the third season of the year. Updated with every battle this season, data is reseted at the end of the year..
RegularTeam *RegularTeam`json:"regular_team" xml:"regular_team" ` // Battle statistics of permanent teams
StrongholdDefense *StrongholdDefense`json:"stronghold_defense" xml:"stronghold_defense" ` // General stats for player's battles in Stronghold defense
StrongholdSkirmish *StrongholdSkirmish`json:"stronghold_skirmish" xml:"stronghold_skirmish" ` // General stats for player's battles in Skirmishes
Team *Team`json:"team" xml:"team" ` // Team battles statistics
}