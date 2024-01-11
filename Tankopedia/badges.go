package Tankopedia

import (
	"encoding/json"

	wgapi "github.com/Thenecromance/WarGamingAPI"
	structure "github.com/Thenecromance/WarGamingAPI/Tankopedia/structure/badges"
)

type badgesResponse struct {
	wgapi.ResponseBase
	Data map[string]structure.Badges `json:"data"`
}

func Badges() map[string]structure.Badges {
	request, err := wgapi.Request(buildPath(badges), map[string]string{"language": "zh-cn"})
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
