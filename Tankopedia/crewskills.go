package Tankopedia

import (
	"encoding/json"

	wgapi "github.com/Thenecromance/WGAPI"
	structure "github.com/Thenecromance/WGAPI/Tankopedia/structure/crewskills"
)

type ImageUrl struct {
	BigIcon   string `json:"big_icon"`
	SmallIcon string `json:"small_icon"`
}

type crewskillsResponse struct {
	wgapi.ResponseBase
	Data map[string]structure.Crewskills `json:"data"`
}

func CrewSkills() map[string]structure.Crewskills {
	request, err := wgapi.Request(buildPath(crewskills), map[string]string{"language": "zh-cn"})
	if err != nil {
		wgapi.Logger.Error(err)
		return nil
	}

	//parse the json
	resp := &crewskillsResponse{}
	jsonErr := json.Unmarshal(request, resp)
	if jsonErr != nil {
		wgapi.Logger.Error(jsonErr)
		return nil
	}
	return resp.Data
}
