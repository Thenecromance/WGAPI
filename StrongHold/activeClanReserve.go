package StrongHold

import (
	"encoding/json"

	wgapi "github.com/Thenecromance/WarGamingAPI"
)

type activeClanReservesResponse struct {
	wgapi.ResponseBase
	ActivatedAt string `json:"data"`
}

func ActiveClanReserves(access_token string, reserve_level string, reserve_type string) string {

	panic("Don't have time to implement this")
	request, err := wgapi.Request(buildPath(activateclanreserve), map[string]string{
		"access_token":  access_token,
		"reserve_level": reserve_level,
		"reserve_type":  reserve_type,
	})

	if err != nil {
		wgapi.Logger.Error(err)
		return ""
	}

	//parse the json
	resp := &activeClanReservesResponse{}
	jsonErr := json.Unmarshal(request, resp)
	if jsonErr != nil {
		wgapi.Logger.Error(jsonErr)
		return ""
	}

	return resp.ActivatedAt
}
