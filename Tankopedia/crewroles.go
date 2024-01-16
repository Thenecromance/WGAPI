package Tankopedia

import (
	"encoding/json"

	wgapi "github.com/Thenecromance/WGAPI"
	structure "github.com/Thenecromance/WGAPI/Tankopedia/structure/crewroles"
)

type crewrolesResponse struct {
	wgapi.ResponseBase
	Data map[string]structure.Crewroles `json:"data"`
}

func CrewRoles(opts ...wgapi.ParamOption) map[string]structure.Crewroles {
	wgapi.InsertBefore(opts,
		wgapi.WithParam("language", "zh-cn"),
		wgapi.WithPath(buildPath(crewroles)),
	)
	request, err := wgapi.Request(opts...)
	
	if err != nil {
		wgapi.Logger.Error(err)
		return nil
	}

	//parse the json
	resp := &crewrolesResponse{}
	jsonErr := json.Unmarshal(request, resp)
	if jsonErr != nil {
		wgapi.Logger.Error(jsonErr)
		return nil
	}
	return resp.Data
}
