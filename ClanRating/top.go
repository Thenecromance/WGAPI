package ClanRating

import (
	"encoding/json"

	wgapi "github.com/Thenecromance/WGAPI"
	structure "github.com/Thenecromance/WGAPI/ClanRating/structure/top"
)

type topResponse struct {
	wgapi.ResponseBase
	Data []structure.Top `json:"data"`
}

func Top(rank_field string, opts ...wgapi.ParamOption) []structure.Top {
	wgapi.InsertBefore(opts,
		wgapi.WithParam("rank_field", rank_field),
		wgapi.WithParam("language", "zh-cn"),
		wgapi.WithParam("limit", "100"),
		wgapi.WithPath(buildPath(top)),
	)
	request, err := wgapi.Request(opts...)
	
	if err != nil {
		wgapi.Logger.Error(err)
		return nil
	}

	//parse the json
	resp := &topResponse{}
	jsonErr := json.Unmarshal(request, resp)
	if jsonErr != nil {
		wgapi.Logger.Error(jsonErr)
		return nil
	}
	return resp.Data
}
