package ClanRating

import (
	"encoding/json"

	wgapi "github.com/Thenecromance/WarGamingAPI"
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

func TypeOfRatings() RankFields {
	request, err := wgapi.Request(buildPath(types), map[string]string{"language": "zh-cn"})
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
