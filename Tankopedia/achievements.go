package Tankopedia

import (
	"encoding/json"
	//
	wgapi "github.com/Thenecromance/WarGamingAPI"
	structure "github.com/Thenecromance/WarGamingAPI/Tankopedia/structure/achievements"
)

type AData map[string]structure.Achievements

type achievementsResponse struct {
	wgapi.ResponseBase
	Data AData `json:"data"`
}

func Achievements() map[string]structure.Achievements {
	request, err := wgapi.Request(buildPath(achievements), map[string]string{"language": "zh-cn"})
	if err != nil {
		wgapi.Logger.Error(err)
		return nil
	}

	//parse the json
	resp := &achievementsResponse{}
	jsonErr := json.Unmarshal(request, resp)
	if jsonErr != nil {
		wgapi.Logger.Error(jsonErr)
		return nil
	}
	return resp.Data
}
