package Account

import (
	"encoding/json"
	"fmt"

	wgapi "github.com/Thenecromance/WGAPI"
)

type Player struct {
	AccountId int64  `json:"account_id"`
	NickName  string `json:"nickname"`
}

type responsePlayer struct {
	wgapi.ResponseBase
	Data []Player `json:"data"`
}

func /*  (this *Service) */ Players(name string, opts ...wgapi.ParamOption) (ret []Player) {
	wgapi.InsertBefore(opts, wgapi.WithParam("search", name), wgapi.WithPath(buildPath(players)))

	request, err := wgapi.Request(opts...)

	if err != nil {
		return nil
	}

	//parse the json
	resp := &responsePlayer{}
	jsonErr := json.Unmarshal(request, resp)
	if jsonErr != nil {
		fmt.Println(jsonErr)
		return nil
	}
	return resp.Data
}
