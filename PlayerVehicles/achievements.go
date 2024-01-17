package PlayerVehicles

import (
	"encoding/json"

	wgapi "github.com/Thenecromance/WGAPI"
	structure "github.com/Thenecromance/WGAPI/PlayerVehicles/structure/achievements"
)

// type Data map[string]interface{}

type Data = map[string][]structure.Achievements
type achivementResponse struct {
	wgapi.ResponseBase
	Data Data `json:"data"`
}

func Achievements(account_id string, opts ...wgapi.ParamOption) map[string][]structure.Achievements {
	wgapi.InsertBefore(&opts,
		wgapi.WithParam("account_id", account_id),
		wgapi.WithParam("language", "zh-cn"),
		wgapi.WithPath(buildPath(achievements)),
	)
	request, err := wgapi.Request(opts...)

	if err != nil {
		wgapi.Logger.Error(err)
		return nil
	}

	//parse the json
	resp := &achivementResponse{}
	jsonErr := json.Unmarshal(request, resp)

	if jsonErr != nil {
		wgapi.Logger.Error(jsonErr)
		return nil
	}

	if resp.Status != "ok" {
		wgapi.Logger.Error("resp.Status is ", resp.Status)
		return nil
	}

	return resp.Data
}
