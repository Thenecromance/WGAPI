package Tankopedia

import (
	"encoding/json"

	wgapi "github.com/Thenecromance/WGAPI"
	structure "github.com/Thenecromance/WGAPI/Tankopedia/structure/info"
)

type AchievementSections struct {
	Name  string `json:"name"`
	Order int    `json:"order"`
}
type informationResponse struct {
	wgapi.ResponseBase
	Data structure.Info `json:"data"`
}

func Information(opts ...wgapi.ParamOption) structure.Info {

	wgapi.InsertBefore(opts,
		wgapi.WithParam("language", "zh-cn"),
		wgapi.WithPath(buildPath(info)),
	)
	request, err := wgapi.Request(opts...)

	if err != nil {
		wgapi.Logger.Error(err)
		return structure.Info{}
	}

	//parse the json
	resp := &informationResponse{}
	jsonErr := json.Unmarshal(request, resp)
	wgapi.Logger.Debugf("%s", resp.Status)
	if jsonErr != nil {
		wgapi.Logger.Error(jsonErr)
		return structure.Info{}
	}
	return resp.Data
}
