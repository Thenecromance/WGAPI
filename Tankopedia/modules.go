package Tankopedia

import (
	"encoding/json"

	wgapi "github.com/Thenecromance/WarGamingAPI"
	structure "github.com/Thenecromance/WarGamingAPI/Tankopedia/structure/modules"
)

type modulesResponse struct {
	wgapi.ResponseBase
	Data map[string]structure.Modules `json:"data"`
}

func Modules() map[string]structure.Modules {
	request, err := wgapi.Request(buildPath(modules), map[string]string{"extra": "default_profile", "fields": "default_profile.engine"})
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
