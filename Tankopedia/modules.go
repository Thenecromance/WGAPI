package Tankopedia

import (
	"encoding/json"

	wgapi "github.com/Thenecromance/WGAPI"
	structure "github.com/Thenecromance/WGAPI/Tankopedia/structure/modules"
)

type modulesResponse struct {
	wgapi.ResponseBase
	Data map[string]structure.Modules `json:"data"`
}

func Modules(opts ...wgapi.ParamOption) map[string]structure.Modules {

	wgapi.InsertBefore(opts,
		wgapi.WithParam("extra", "default_profile"),
		wgapi.WithParam("fields", "default_profile.engine"),
		wgapi.WithParam("language", "zh-cn"),
		wgapi.WithPath(buildPath(info)),
	)
	request, err := wgapi.Request(opts...)

	if err != nil {
		wgapi.Logger.Error(err)
		return nil
	}

	//parse the json
	resp := &modulesResponse{}
	jsonErr := json.Unmarshal(request, resp)
	if jsonErr != nil {
		wgapi.Logger.Error(jsonErr)
		return nil
	}
	return resp.Data
}
