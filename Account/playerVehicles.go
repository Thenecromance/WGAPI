package Account

import (
	"encoding/json"
	"fmt"

	wgapi "github.com/Thenecromance/WGAPI"
)

// Player's Vehicles Information
type Vehicle struct {
	MarkOfMastery     int64 `json:"mark_of_mastery"`
	TankId            int64 `json:"tank_id"`
	VehicleStatistics `json:"statistics"`
}

type VehicleStatistics struct {
	Battles int64 `json:"battles"`
	Wins    int64 `json:"wins"`
}

type VehData map[string][]Vehicle

type playerVehiclesResponse struct {
	wgapi.ResponseBase
	VehData `json:"data"`
}

func /* (this *Service) */ PlayerVehicles(accountIds string, opts ...wgapi.ParamOption) []PlayerDetail {

	wgapi.InsertBefore(&opts,
		wgapi.WithParam("account_id", accountIds),
		wgapi.WithPath(buildPath(vehicles)),
	)
	request, err := wgapi.Request(opts...)
	if err != nil {
		return nil
	}
	resp := &playerVehiclesResponse{}
	jsonErr := json.Unmarshal(request, resp)
	if jsonErr != nil {
		fmt.Println(jsonErr)
		return nil
	}
	fmt.Println(resp)
	return nil
}
