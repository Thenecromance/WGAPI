package Tankopedia

import (
	"encoding/json"

	wgapi "github.com/Thenecromance/WGAPI"
	structure "github.com/Thenecromance/WGAPI/Tankopedia/structure/vehicles"
)

type Data map[string]structure.Vehicles
type vehiclesResponse struct {
	wgapi.ResponseBase
	Data Data `json:"data"`
}

func Vehicles(tank_id string) map[string]structure.Vehicles {
	request, err := wgapi.Request(buildPath(vehicles), map[string]string{"tank_id": tank_id})
	if err != nil {

		if wgapi.Logger != nil {
			wgapi.Logger.Error(err)
		}

		return nil
	}

	//parse the json
	resp := &vehiclesResponse{}
	jsonErr := json.Unmarshal(request, resp)
	if jsonErr != nil {
		wgapi.Logger.Error(jsonErr)
		return nil
	}
	return resp.Data
}
