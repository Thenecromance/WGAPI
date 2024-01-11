package Tankopedia

import (
	"encoding/json"

	wgapi "github.com/Thenecromance/WGAPI"
	structure "github.com/Thenecromance/WGAPI/Tankopedia/structure/vehicleprofile"
)

type VpData map[string]structure.Vehicleprofile
type vehicleProfileResponse struct {
	wgapi.ResponseBase
	Data VpData `json:"data"`
}

func VehicleCharacteristics(tank_id string) structure.Vehicleprofile {
	request, err := wgapi.Request(buildPath(vehicleprofile), map[string]string{"tank_id": tank_id})
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
