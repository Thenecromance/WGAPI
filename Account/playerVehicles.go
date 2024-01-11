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

func /* (this *Service) */ PlayerVehicles(accountIds string) []PlayerDetail {

	request, err := wgapi.Request(buildPath(vehicles) /* +this.methodPath["vehicles"] */, map[string]string{"account_id": accountIds})
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
