package PlayerVehicles

import (
	"encoding/json"

	wgapi "github.com/Thenecromance/WGAPI"
	structure "github.com/Thenecromance/WGAPI/PlayerVehicles/structure/mastery"
)

// type Data map[string]interface{}

type MasteryData struct {
	All map[string][]structure.Mastery `json:"all"`
}
type masteryResponse struct {
	wgapi.ResponseBase
	Data structure.Mastery `json:"data"`
}

func MasteryDatas(distribution string, percentile string) structure.Mastery {
	request, err := wgapi.Request(buildPath(mastery), map[string]string{"distribution": distribution, "percentile": percentile, "language": "zh-cn"})
	if err != nil {
		wgapi.Logger.Error(err)
		return structure.Mastery{}
	}

	//parse the json
	resp := &masteryResponse{}
	jsonErr := json.Unmarshal(request, resp)

	if jsonErr != nil {
		wgapi.Logger.Error(jsonErr)
		return structure.Mastery{}
	}

	if resp.Status != "ok" {
		wgapi.Logger.Error("resp.Status is ", resp.Status)
		return structure.Mastery{}
	}

	return resp.Data
}
