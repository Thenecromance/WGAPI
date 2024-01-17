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

func Arenas(opts ...wgapi.ParamOption) map[string]Map {
	wgapi.InsertBefore(&opts,
		wgapi.WithParam("language", "zh-cn"),
		wgapi.WithPath(buildPath(arenas)),
	)
	request, err := wgapi.Request(opts...)
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
