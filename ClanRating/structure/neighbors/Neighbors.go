package structure
type Neighbors struct {
ClanId int32/*numeric*/ `json:"clan_id" xml:"clan_id" ` // Clan ID

ClanName string/*string*/ `json:"clan_name" xml:"clan_name" ` // Clan name

ClanTag string/*string*/ `json:"clan_tag" xml:"clan_tag" ` // Clan tag

ExcludeReasons map[string]string/*associative array*/ `json:"exclude_reasons" xml:"exclude_reasons" ` // Reasons why specified rating categories were not calculated. Contains data in "key-value" format, where the key is category name and the value is reason.

BattlesCountAvg *BattlesCountAvg`json:"battles_count_avg" xml:"battles_count_avg" ` // Average number of battles
BattlesCountAvgDaily *BattlesCountAvgDaily`json:"battles_count_avg_daily" xml:"battles_count_avg_daily" ` // Average number of battles per day
Efficiency *Efficiency`json:"efficiency" xml:"efficiency" ` // Indicator of clan's performance.
FbEloRating *FbEloRating`json:"fb_elo_rating" xml:"fb_elo_rating" ` // Weighted Elo rating achieved in the Stronghold mode
FbEloRating10 *FbEloRating10`json:"fb_elo_rating_10" xml:"fb_elo_rating_10" ` // Elo rating achieved by the clan on Tier X vehicles in the Stronghold mode
FbEloRating6 *FbEloRating6`json:"fb_elo_rating_6" xml:"fb_elo_rating_6" ` // Elo rating achieved on Tier VI vehicles in the Stronghold mode
FbEloRating8 *FbEloRating8`json:"fb_elo_rating_8" xml:"fb_elo_rating_8" ` // Elo rating achieved on Tier VIII vehicles in the Stronghold mode
GlobalRatingAvg *GlobalRatingAvg`json:"global_rating_avg" xml:"global_rating_avg" ` // Average global rating value
GlobalRatingWeightedAvg *GlobalRatingWeightedAvg`json:"global_rating_weighted_avg" xml:"global_rating_weighted_avg" ` // Weighted average of global rating value
GmEloRating *GmEloRating`json:"gm_elo_rating" xml:"gm_elo_rating" ` // Elo rating on the Global Map
GmEloRating10 *GmEloRating10`json:"gm_elo_rating_10" xml:"gm_elo_rating_10" ` // Elo rating on the Global Map in Absolute division
GmEloRating6 *GmEloRating6`json:"gm_elo_rating_6" xml:"gm_elo_rating_6" ` // Elo rating on the Global Map in Medium division
GmEloRating8 *GmEloRating8`json:"gm_elo_rating_8" xml:"gm_elo_rating_8" ` // Elo rating on the Global Map in Champion division
RatingFort *RatingFort`json:"rating_fort" xml:"rating_fort" ` // Rating in Battles for Stronghold
V10lAvg *V10lAvg`json:"v10l_avg" xml:"v10l_avg" ` // Average number of vehicles of Tier 10 per clan member
WinsRatioAvg *WinsRatioAvg`json:"wins_ratio_avg" xml:"wins_ratio_avg" ` // Average victory rate
}