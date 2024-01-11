package Tankopedia

import (
	"encoding/json"

	wgapi "github.com/Thenecromance/WarGamingAPI"
	structure "github.com/Thenecromance/WarGamingAPI/Tankopedia/structure/boosters"
)

type Images struct {
	Large string `json:"large"`
	Small string `json:"small"`
}

type boostersResponse struct {
	wgapi.ResponseBase
	Data map[string]structure.Boosters `json:"data"`
}

func Boosters() map[string]structure.Boosters {
	request, err := wgapi.Request(buildPath(boosters), map[string]string{"language": "zh-cn"})
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
