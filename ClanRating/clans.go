package ClanRating

import (
	"encoding/json"

	wgapi "github.com/Thenecromance/WGAPI"
)

type BattlesCountBase struct {
	Rank      int32   `json:"rank"`
	RankDelta int32   `json:"rank_delta"`
	Value     float64 `json:"value"`
}

type BattlesCountAvg = BattlesCountBase
type BattlesCountAvgDaily = BattlesCountBase
type Efficiency = BattlesCountBase
type FbEloRating = BattlesCountBase
type FbEloRating10 = BattlesCountBase
type FbEloRating6 = BattlesCountBase
type FbEloRating8 = BattlesCountBase
type GlobalRatingAvg = BattlesCountBase
type GlobalRatingWeightedAvg = BattlesCountBase
type GmEloRating = BattlesCountBase
type GmEloRating10 = BattlesCountBase
type GmEloRating6 = BattlesCountBase
type GmEloRating8 = BattlesCountBase
type RatingFort = BattlesCountBase
type V10lAvg = BattlesCountBase
type WinsRatioAvg = BattlesCountBase

type ClanRating struct {
	ClanId                  int32                   `json:"clan_id"`
	ClanName                string                  `json:"clan_name"`
	ClanTag                 string                  `json:"clan_tag"`
	ExcludeReasons          map[string]string       `json:"exclude_reasons"`
	BattlesCountAvg         BattlesCountAvg         `json:"battles_count_avg"`
	BattlesCountAvgDaily    BattlesCountAvgDaily    `json:"battles_count_avg_daily"`
	Efficiency              Efficiency              `json:"efficiency"`
	FbEloRating             FbEloRating             `json:"fb_elo_rating"`
	FbEloRating6            FbEloRating6            `json:"fb_elo_rating_6"`
	FbEloRating8            FbEloRating8            `json:"fb_elo_rating_8"`
	FbEloRating10           FbEloRating10           `json:"fb_elo_rating_10"`
	GlobalRatingAvg         GlobalRatingAvg         `json:"global_rating_avg"`
	GlobalRatingWeightedAvg GlobalRatingWeightedAvg `json:"global_rating_weighted_avg"`
	GmEloRating             GmEloRating             `json:"gm_elo_rating"`
	GmEloRating6            GmEloRating6            `json:"gm_elo_rating_6"`
	GmEloRating8            GmEloRating8            `json:"gm_elo_rating_8"`
	GmEloRating10           GmEloRating10           `json:"gm_elo_rating_10"`
	RatingFort              RatingFort              `json:"rating_fort"`
	V10lAvg                 V10lAvg                 `json:"v10l_avg"`
	WinsRatioAvg            WinsRatioAvg            `json:"wins_ratio_avg"`
}

type clanRatingsResponse struct {
	wgapi.ResponseBase
	Data map[string]ClanRating `json:"data"`
}

func ClanRatings(clan_id string, opts ...wgapi.ParamOption) map[string]ClanRating {
	wgapi.InsertBefore(&opts,
		wgapi.WithParam("clan_id", clan_id),
		wgapi.WithParam("language", "zh-cn"),
		wgapi.WithPath(buildPath(clans)),
	)
	request, err := wgapi.Request(opts...)
	if err != nil {
		wgapi.Logger.Error(err)
		return nil
	}

	//parse the json
	resp := &clanRatingsResponse{}
	jsonErr := json.Unmarshal(request, resp)
	if jsonErr != nil {
		wgapi.Logger.Error(jsonErr)
		return nil
	}
	return resp.Data
}
