package Tankopedia

import (
	"encoding/json"

	wgapi "github.com/Thenecromance/WGAPI"
)

type Map struct {
	ArenaId        string `json:"arena_id"`
	CamouflageType string `json:"camouflage_type"`
	Description    string `json:"description"`
	NameI18n       string `json:"name_i18n"`
}

type mapResponse struct {
	wgapi.ResponseBase
	Data map[string]Map `json:"data"`
}

func Arenas() map[string]Map {
	request, err := wgapi.Request(buildPath(arenas), map[string]string{"language": "zh-cn"})
	if err != nil {
		wgapi.Logger.Error(err)
		return nil
	}

	//parse the json
	resp := &mapResponse{}
	jsonErr := json.Unmarshal(request, resp)
	if jsonErr != nil {
		wgapi.Logger.Error(jsonErr)
		return nil
	}
	return resp.Data
}
