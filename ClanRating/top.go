package ClanRating

import (
	"encoding/json"

	wgapi "github.com/Thenecromance/WarGamingAPI"
	structure "github.com/Thenecromance/WarGamingAPI/ClanRating/structure/top"
)

type topResponse struct {
	wgapi.ResponseBase
	Data []structure.Top `json:"data"`
}

func Top(rank_field string) []structure.Top {
	request, err := wgapi.Request(buildPath(top), map[string]string{"rank_field": rank_field, "language": "zh-cn", "limit": "100"})
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
