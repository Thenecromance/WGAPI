package ClanRating

import (
	"encoding/json"

	wgapi "github.com/Thenecromance/WarGamingAPI"
)

type neighborsResponse struct {
	wgapi.ResponseBase
	Data []ClanRating `json:"data"`
}

func Neighbors(clan_id string, rank_field string) []ClanRating {
	request, err := wgapi.Request(buildPath(neighbors), map[string]string{"clan_id": clan_id, "rank_field": rank_field, "language": "zh-cn"})
	if err != nil {
		wgapi.Logger.Error(err)
		return nil
	}

	//parse the json
	resp := &neighborsResponse{}
	jsonErr := json.Unmarshal(request, resp)
	if jsonErr != nil {
		wgapi.Logger.Error(jsonErr)
		return nil
	}
	return resp.Data
}
