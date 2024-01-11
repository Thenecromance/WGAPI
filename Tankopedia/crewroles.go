package Tankopedia

import (
	"encoding/json"

	wgapi "github.com/Thenecromance/WarGamingAPI"
	structure "github.com/Thenecromance/WarGamingAPI/Tankopedia/structure/crewroles"
)

type crewrolesResponse struct {
	wgapi.ResponseBase
	Data map[string]structure.Crewroles `json:"data"`
}

func CrewRoles() map[string]structure.Crewroles {
	request, err := wgapi.Request(buildPath(crewroles), map[string]string{"language": "zh-cn"})
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
