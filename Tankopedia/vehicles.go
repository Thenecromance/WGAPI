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

// TODO: something goes wrong here , seems like parse failed
func Vehicles(tank_id string, opts ...wgapi.ParamOption) structure.Vehicles {

	wgapi.InsertBefore(&opts,
		wgapi.WithParam("tank_id", tank_id),
		wgapi.WithParam("language", "zh-cn"),
		wgapi.WithPath(buildPath(vehicles)),
	)
	request, err := wgapi.Request(opts...)

	//request, err := wgapi.Request(buildPath(param.GetRegion(), vehicles), map[string]string{"tank_id": tank_id})
	if err != nil {

		if wgapi.Logger != nil {
			wgapi.Logger.Error(err)
		}

		return structure.Vehicles{}
	}

	//parse the json
	resp := &vehiclesResponse{}
	jsonErr := json.Unmarshal(request, resp)
	if jsonErr != nil {
		wgapi.Logger.Error(jsonErr)
		return structure.Vehicles{}
	}
	return resp.Data[tank_id]
}

/*
func AllVehicles() map[string]structure.Vehicles {
	request, err := wgapi.Request(buildPath(param.GetRegion(), vehicles), map[string]string{
		"language": "zh-tw",
	})
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
*/
