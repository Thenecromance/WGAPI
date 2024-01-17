package StrongHold

import (
	"encoding/json"

	wgapi "github.com/Thenecromance/WGAPI"
)

type activeClanReservesResponse struct {
	wgapi.ResponseBase
	ActivatedAt string `json:"data"`
}

func ActiveClanReserves(access_token string, reserve_level string, reserve_type string, opts ...wgapi.ParamOption) string {

	panic("Don't have time to implement this")

	wgapi.InsertBefore(&opts,
		wgapi.WithParam("access_token", access_token),
		wgapi.WithParam("reserve_level", reserve_level),
		wgapi.WithParam("reserve_type", reserve_type),
		wgapi.WithPath(buildPath(activateclanreserve)),
	)
	request, err := wgapi.Request(opts...)

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
