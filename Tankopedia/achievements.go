package Tankopedia

import (
	"encoding/json"
	//
	wgapi "github.com/Thenecromance/WGAPI"
	structure "github.com/Thenecromance/WGAPI/Tankopedia/structure/achievements"
)

type AData map[string]structure.Achievements

type achievementsResponse struct {
	wgapi.ResponseBase
	Data AData `json:"data"`
}

func Achievements(opts ...wgapi.ParamOption) map[string]structure.Achievements {
	wgapi.InsertBefore(opts,
		wgapi.WithParam("language", "zh-cn"),
		wgapi.WithPath(buildPath(achievements)),
	)
	request, err := wgapi.Request(opts...)
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
