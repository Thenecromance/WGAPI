package ClanRating

import (
	"encoding/json"

	wgapi "github.com/Thenecromance/WGAPI"
)

type RankFields struct {
	Type       string   `json:"type"`
	RankFields []string `json:"rank_fields"`
}
type Data map[string]RankFields
type rankFieldsResponse struct {
	wgapi.ResponseBase
	Data `json:"data"`
}

func TypeOfRatings(opts ...wgapi.ParamOption) RankFields {
	wgapi.InsertBefore(opts,
		wgapi.WithParam("language", "zh-cn"),
		wgapi.WithPath(buildPath(types)),
	)
	request, err := wgapi.Request(opts...)
	
	if err != nil {
		wgapi.Logger.Error(err)
		return RankFields{}
	}

	//parse the json
	resp := &rankFieldsResponse{}
	jsonErr := json.Unmarshal(request, resp)
	if jsonErr != nil {
		wgapi.Logger.Error(jsonErr)
		return RankFields{}
	}
	return resp.Data["all"]
}
