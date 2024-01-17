package Tankopedia

import (
	"encoding/json"

	wgapi "github.com/Thenecromance/WGAPI"
	structure "github.com/Thenecromance/WGAPI/Tankopedia/structure/vehicleprofiles"
)

/*
	 type VehicleProfiles struct {
		IsDefault   bool   `json:"is_default"`
		PriceCredit int64  `json:"price_credit"`
		ProfileId   string `json:"profile_id"`
		TankId      int64  `json:"tank_id"`
	}
*/
type vehicleprofilesResponse struct {
	wgapi.ResponseBase
	Data map[string][]structure.Vehicleprofiles `json:"data"`
}

func VehicleProfiles(tank_id string, opts ...wgapi.ParamOption) map[string][]structure.Vehicleprofiles {
	wgapi.InsertBefore(&opts,
		wgapi.WithParam("language", "zh-cn"),
		wgapi.WithParam("tank_id", tank_id),
		wgapi.WithPath(buildPath(vehicleprofiles)),
	)
	request, err := wgapi.Request(opts...)

	if err != nil {
		wgapi.Logger.Error(err)
		return nil
	}

	//parse the json
	resp := &vehicleprofilesResponse{}
	jsonErr := json.Unmarshal(request, resp)
	if jsonErr != nil {
		wgapi.Logger.Error(jsonErr)
		return nil
	}
	return resp.Data
}
