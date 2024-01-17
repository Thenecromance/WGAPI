package wargaming

import (
	"encoding/json"
	"fmt"

	wgapi "github.com/Thenecromance/WGAPI"
)

type ServerStatus struct {
	PlayersOnline int    `json:"players_online"`
	Server        string `json:"server"`
}

type achievementsResponse struct {
	wgapi.ResponseBase
	Data map[string][]ServerStatus `json:"data"`
}

func GameServers(opts ...wgapi.ParamOption) map[string][]ServerStatus {

	/*opts =*/ wgapi.InsertBefore(&opts,
		wgapi.WithParam("language", "zh-cn"),
		wgapi.WithPath(buildPath(info)),
	)
	request, err := wgapi.Request(opts...)

	if err != nil {
		fmt.Println(err)
		wgapi.Logger.Error(err)
		return nil
	}

	//parse the json
	resp := &achievementsResponse{}
	jsonErr := json.Unmarshal(request, resp)
	if jsonErr != nil {
		fmt.Println(jsonErr)
		wgapi.Logger.Error(jsonErr)
		return nil
	}
	return resp.Data

}
