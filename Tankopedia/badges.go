package Tankopedia

import (
	"encoding/json"

	wgapi "github.com/Thenecromance/WGAPI"
	structure "github.com/Thenecromance/WGAPI/Tankopedia/structure/badges"
)

type badgesResponse struct {
	wgapi.ResponseBase
	Data map[string]structure.Badges `json:"data"`
}

func Badges(opts ...wgapi.ParamOption) map[string]structure.Badges {
	wgapi.InsertBefore(opts,
		wgapi.WithParam("language", "zh-cn"),
		wgapi.WithPath(buildPath(badges)),
	)
	request, err := wgapi.Request(opts...)
	if err != nil {
		wgapi.Logger.Error(err)
		return nil
	}

	//parse the json
	resp := &badgesResponse{}
	jsonErr := json.Unmarshal(request, resp)
	if jsonErr != nil {
		wgapi.Logger.Error(jsonErr)
		return nil
	}
	return resp.Data
}
