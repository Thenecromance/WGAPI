package Tankopedia

import (
	"encoding/json"

	wgapi "github.com/Thenecromance/WarGamingAPI"
	structure "github.com/Thenecromance/WarGamingAPI/Tankopedia/structure/personalmissions"
)

//	type PersonalMissions struct {
//		CampaignId  int64                `json:"campaign_id"`
//		Description string               `json:"description"`
//		Name        string               `json:"name"`
//		Operations  map[string]Operation `json:"operations"`
//	}
//
//	type Missions struct {
//		Description string             `json:"description"`
//		Hint        string             `json:"hint"`
//		MaxLevel    int64              `json:"max_level"`
//		MinLevel    int64              `json:"min_level"`
//		MissionId   int64              `json:"mission_id"`
//		Name        string             `json:"name"`
//		SetId       int64              `json:"set_id"`
//		Tags        []string           `json:"tags"`
//		Rewards     map[string]Rewards `json:"rewards"`
//	}
//
//	type Operation struct {
//		Description   string              `json:"description"`
//		Image         string              `json:"image"`
//		MissionsInSet int64               `json:"missions_in_set"`
//		Name          string              `json:"name"`
//		NextId        int64               `json:"next_id"`
//		OperationId   int64               `json:"operation_id"`
//		SetsCount     int64               `json:"sets_count"`
//		SetsToNext    int64               `json:"sets_to_next"`
//		Missions      map[string]Missions `json:"missions"`
//		Reward        Reward              `json:"reward"`
//	}
//
//	type Rewards struct {
//		Berths     int64            `json:"berths"`
//		Conditions string           `json:"conditions"`
//		Credits    int64            `json:"credits"`
//		FreeXp     int64            `json:"free_xp"`
//		Items      map[string]int64 `json:"items"`
//		Premium    int64            `json:"premium"`
//		Slots      int64            `json:"slots"`
//		Tokens     int64            `json:"tokens"`
//	}
//
//	type Reward struct {
//		Slots int64   `json:"slots"`
//		Tanks []int64 `json:"tanks"`
//	}
type personalmissionsResponse struct {
	wgapi.ResponseBase
	Data map[string]structure.Personalmissions `json:"data"`
}

func PersonalMissions() map[string]structure.Personalmissions {
	request, err := wgapi.Request(buildPath(personalmissions), map[string]string{"language": "zh-cn"})
	if err != nil {
		wgapi.Logger.Error(err)
		return nil
	}

	//parse the json
	resp := &personalmissionsResponse{}
	jsonErr := json.Unmarshal(request, resp)
	if jsonErr != nil {
		wgapi.Logger.Error(jsonErr)
		return nil
	}
	return resp.Data
}
