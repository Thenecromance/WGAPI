package Clans

import (
	"encoding/json"

	wgapi "github.com/Thenecromance/WarGamingAPI"
)

type Data map[string]interface{}
type clanGlossaryResponse struct {
	wgapi.ResponseBase
	Data Data `json:"data"`
}

func Glossary() map[string]string {
	request, err := wgapi.Request(buildPath(glossary), map[string]string{"language": "zh-cn"})
	if err != nil {
		wgapi.Logger.Error(err)
		return nil
	}

	//parse the json
	resp := &clanGlossaryResponse{}
	jsonErr := json.Unmarshal(request, resp)
	if jsonErr != nil {
		wgapi.Logger.Error(jsonErr)
		return nil
	}
	//fmt.Println(resp)
	//TODO: not implemented
	panic("not implemented")
	return nil
	//return resp.Data
	//return resp.Data["clan_roles"]
}