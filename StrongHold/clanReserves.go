package StrongHold

import (
	"encoding/json"

	wgapi "github.com/Thenecromance/WGAPI"
	structure "github.com/Thenecromance/WGAPI/StrongHold/structure/clanreserves"
)

type clanReservesResponse struct {
	wgapi.ResponseBase
	Data []structure.Clanreserves `json:"data"`
}

func ClanReserves(access_token string, opts ...wgapi.ParamOption) []structure.Clanreserves {

	wgapi.InsertBefore(&opts,
		wgapi.WithParam("access_token", access_token),
		wgapi.WithPath(buildPath(clanreserves)),
	)
	request, err := wgapi.Request(opts...)

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
