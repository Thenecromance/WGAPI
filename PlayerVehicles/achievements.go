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

func Achievements(account_id string) map[string][]structure.Achievements {
	request, err := wgapi.Request(buildPath(achievements), map[string]string{"account_id": account_id, "language": "zh-cn"})
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
