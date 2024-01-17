package ClanRating

import (
	"encoding/json"

	wgapi "github.com/Thenecromance/WGAPI"
	"github.com/Thenecromance/WGAPI/ClanRating/structure"
)

type DatesWithAvailableRatingsResponse struct {
	wgapi.ResponseBase
	Data map[string]structure.DatesWithAvailableRatings `json:"data"`
}

func DatesWithAvailableRatings(opts ...wgapi.ParamOption) structure.DatesWithAvailableRatings {
	wgapi.InsertBefore(&opts,
		wgapi.WithParam("language", "zh-cn"),
		wgapi.WithPath(buildPath(dates)),
	)
	request, err := wgapi.Request(opts...)
	if err != nil {
		wgapi.Logger.Error(err)
		return structure.DatesWithAvailableRatings{}
	}

	//parse the json
	resp := &DatesWithAvailableRatingsResponse{}
	jsonErr := json.Unmarshal(request, resp)
	if jsonErr != nil {
		wgapi.Logger.Error(jsonErr)
		return structure.DatesWithAvailableRatings{}
	}
	return resp.Data["all"]
}
