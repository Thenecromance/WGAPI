package PlayerVehicles

import (
	"encoding/json"

	wgapi "github.com/Thenecromance/WarGamingAPI"
	structure "github.com/Thenecromance/WarGamingAPI/PlayerVehicles/structure/stats"
)

type statsData map[string]structure.Stats
type StatsRatingsResponse struct {
	wgapi.ResponseBase
	Data statsData `json:"data"`
}

func Stats(account_id string) structure.Stats {
	request, err := wgapi.Request(buildPath(stats), map[string]string{"account_id": account_id, "language": "zh-cn"})
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
