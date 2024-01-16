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

func GameServers() map[string][]ServerStatus {
	request, err := wgapi.Request(buildPath(info), map[string]string{"language": "zh-cn"})
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
