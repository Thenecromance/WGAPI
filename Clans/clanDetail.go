package Clans

import (
	"encoding/json"

	wgapi "github.com/Thenecromance/WGAPI"

	structure "github.com/Thenecromance/WGAPI/Clans/structure/ClanDetails"
)

type timestamp = int64

// type ClanDetail struct {
// 	AcceptsJoinRequests bool                `json:"accepts_join_requests"`
// 	ClanId              int64               `json:"clan_id"`
// 	Color               string              `json:"color"`
// 	CreatedAt           int64               `json:"created_at"`
// 	CreatorId           int64               `json:"creator_id"`
// 	CreatorName         string              `json:"creator_name"`
// 	Description         string              `json:"description"`
// 	DescriptionHtml     string              `json:"description_html"`
// 	IsClanDisbanded     bool                `json:"is_clan_disbanded"`
// 	LeaderId            int64               `json:"leader_id"`
// 	LeaderName          string              `json:"leader_name"`
// 	MembersCount        int64               `json:"members_count"` //Number of clan members
// 	Motto               string              `json:"motto"`         //Clan motto
// 	Name                string              `json:"name"`          //Clan name
// 	OldName             string              `json:"old_name"`      //Old clan name
// 	OldTag              string              `json:"old_tag"`       //Old clan tag
// 	RenamedAt           timestamp           `json:"renamed_at"`    //Time (UTC) when clan name was changed
// 	Tag                 string              `json:"tag"`           //Clan tag
// 	UpdatedAt           timestamp           `json:"updated_at"`    //Time when clan details were updated
// 	Members             []structure.Members `json:"members"`
// 	structure.Emblems   `json:"emblems"`
// 	*structure.Private  `json:"private"`
// }

type ClanDetails map[string]structure.Clandetails
type ClanDetailResponse struct {
	wgapi.ResponseBase
	Data ClanDetails `json:"data"`
}

func ClanDetail(clanId string, opts ...wgapi.ParamOption) map[string]structure.Clandetails {
	wgapi.InsertBefore(&opts,
		wgapi.WithParam("clan_id", clanId),
		wgapi.WithParam("language", "zh-cn"),
		wgapi.WithPath(buildPath(clan_detail)),
	)
	request, err := wgapi.Request(opts...)

	if err != nil {
		wgapi.Logger.Error(err)
		return nil
	}

	//parse the json
	resp := &ClanDetailResponse{}
	jsonErr := json.Unmarshal(request, resp)
	if jsonErr != nil {
		wgapi.Logger.Error(jsonErr)
		return nil
	}
	return resp.Data
}
