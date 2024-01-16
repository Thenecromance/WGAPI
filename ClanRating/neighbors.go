package ClanRating

import (
	"encoding/json"

	wgapi "github.com/Thenecromance/WGAPI"
)

type neighborsResponse struct {
	wgapi.ResponseBase
	Data []ClanRating `json:"data"`
}

func Neighbors(clanId string, rankField string, opts ...wgapi.ParamOption) []ClanRating {
	wgapi.InsertBefore(opts,
		wgapi.WithParam("clan_id", clanId),
		wgapi.WithParam("rank_field", rankField),
		wgapi.WithParam("language", "zh-cn"),
		wgapi.WithPath(buildPath(neighbors)),
	)
	request, err := wgapi.Request(opts...)
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
