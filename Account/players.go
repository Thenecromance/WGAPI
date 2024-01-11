package Account

import (
	"encoding/json"
	"fmt"

	wgapi "github.com/Thenecromance/WarGamingAPI"
)

type Player struct {
	AccountId int64  `json:"account_id"`
	NickName  string `json:"nickname"`
}

type responsePlayer struct {
	wgapi.ResponseBase
	Data []Player `json:"data"`
}

func /*  (this *Service) */ Players(name string) (ret []Player) {
	request, err := wgapi.Request(buildPath(players) /* +this.methodPath["players"] */, map[string]string{"search": name})
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
