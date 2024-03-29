package Tankopedia

import (
	"encoding/json"

	wgapi "github.com/Thenecromance/WGAPI"
	structure "github.com/Thenecromance/WGAPI/Tankopedia/structure/provisions"
)

/* type Provisions struct {
	Description string `json:"description"`
	Image       string `json:"image"`
	Name        string `json:"name"`
	PriceCredit int64  `json:"price_credit"`
	PriceGold   int64  `json:"price_gold"`
	ProvisionId int64  `json:"provision_id"`
	Tag         string `json:"tag"`
	Type        string `json:"type"`
	Weight      int64  `json:"weight"`
} */

type provisionsResponse struct {
	wgapi.ResponseBase
	Data map[string]structure.Provisions `json:"data"`
}

func Provisions(opts ...wgapi.ParamOption) map[string]structure.Provisions {
	wgapi.InsertBefore(&opts,
		wgapi.WithParam("language", "zh-cn"),
		wgapi.WithPath(buildPath(provisions)),
	)
	request, err := wgapi.Request(opts...)

	if err != nil {
		wgapi.Logger.Error(err)
		return nil
	}

	//parse the json
	resp := &provisionsResponse{}
	jsonErr := json.Unmarshal(request, resp)
	if jsonErr != nil {
		wgapi.Logger.Error(jsonErr)
		return nil
	}
	return resp.Data
}
