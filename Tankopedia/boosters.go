package Tankopedia

import (
	"encoding/json"

	wgapi "github.com/Thenecromance/WGAPI"
	structure "github.com/Thenecromance/WGAPI/Tankopedia/structure/boosters"
)

type Images struct {
	Large string `json:"large"`
	Small string `json:"small"`
}

type boostersResponse struct {
	wgapi.ResponseBase
	Data map[string]structure.Boosters `json:"data"`
}

func Boosters(opts ...wgapi.ParamOption) map[string]structure.Boosters {
	wgapi.InsertBefore(opts,
		wgapi.WithParam("language", "zh-cn"),
		wgapi.WithPath(buildPath(boosters)),
	)
	request, err := wgapi.Request(opts...)
	if err != nil {
		wgapi.Logger.Error(err)
		return nil
	}

	//parse the json
	resp := &boostersResponse{}
	jsonErr := json.Unmarshal(request, resp)
	if jsonErr != nil {
		wgapi.Logger.Error(jsonErr)
		return nil
	}
	return resp.Data
}
