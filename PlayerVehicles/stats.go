package PlayerVehicles

import (
	"encoding/json"

	wgapi "github.com/Thenecromance/WGAPI"
	structure "github.com/Thenecromance/WGAPI/PlayerVehicles/structure/stats"
)

type statsData map[string]structure.Stats
type StatsRatingsResponse struct {
	wgapi.ResponseBase
	Data statsData `json:"data"`
}

func Stats(account_id string, opts ...wgapi.ParamOption) structure.Stats {

	wgapi.InsertBefore(&opts,
		wgapi.WithParam("account_id", account_id),
		wgapi.WithParam("language", "zh-cn"),
		wgapi.WithPath(buildPath(stats)),
	)
	request, err := wgapi.Request(opts...)

	if err != nil {
		wgapi.Logger.Error(err)
		return structure.Stats{}
	}

	//parse the json
	resp := &StatsRatingsResponse{}
	jsonErr := json.Unmarshal(request, resp)
	if jsonErr != nil {
		wgapi.Logger.Error(jsonErr)
		return structure.Stats{}
	}
	return resp.Data["all"]
}
