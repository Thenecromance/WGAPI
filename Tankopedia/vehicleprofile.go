package Tankopedia

import (
	"encoding/json"

	wgapi "github.com/Thenecromance/WGAPI"
	structure "github.com/Thenecromance/WGAPI/Tankopedia/structure/vehicleprofile"
)

type vehicleProfileResponse struct {
	wgapi.ResponseBase
	Data map[string]structure.Vehicleprofile `json:"data"`
}

func VehicleCharacteristics(tank_id string, opts ...wgapi.ParamOption) structure.Vehicleprofile {
	wgapi.InsertBefore(&opts,
		wgapi.WithParam("language", "zh-cn"),
		wgapi.WithParam("tank_id", tank_id),
		wgapi.WithPath(buildPath(vehicleprofile)),
	)
	request, err := wgapi.Request(opts...)

	if err != nil {
		wgapi.Logger.Error(err)
		return structure.Vehicleprofile{}
	}

	//parse the json
	resp := &vehicleProfileResponse{}
	jsonErr := json.Unmarshal(request, resp)
	if jsonErr != nil {
		wgapi.Logger.Error(jsonErr)
		return structure.Vehicleprofile{}
	}
	return resp.Data[tank_id]
}
