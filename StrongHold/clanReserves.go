package StrongHold

import (
	"encoding/json"

	wgapi "github.com/Thenecromance/WarGamingAPI"
	structure "github.com/Thenecromance/WarGamingAPI/StrongHold/structure/clanreserves"
)

type clanReservesResponse struct {
	wgapi.ResponseBase
	Data []structure.Clanreserves `json:"data"`
}

func ClanReserves(access_token string) []structure.Clanreserves {
	request, err := wgapi.Request(buildPath(clanreserves), map[string]string{"access_token": access_token})
	if err != nil {
		wgapi.Logger.Error(err)
		return nil
	}

	//parse the json
	resp := &clanReservesResponse{}
	jsonErr := json.Unmarshal(request, resp)
	if jsonErr != nil {
		wgapi.Logger.Error(jsonErr)
		return nil
	}
	return resp.Data
}
