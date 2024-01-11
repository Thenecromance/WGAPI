package Account

import (
	"encoding/json"
	"fmt"

	wgapi "github.com/Thenecromance/WGAPI"
	structure "github.com/Thenecromance/WGAPI/Account/structure/info"
)

type timestamp = int64

// Player's Basic Information
type PlayerDetail struct {
	AccountId            int64     `json:"account_id" binding:"required"`
	ClanId               int64     `json:"clan_id" binding:"required"`
	CreatedAt            timestamp `json:"created_at" binding:"required"`
	GlobalRating         int64     `json:"global_rating" binding:"required"`
	LastBattleTime       timestamp `json:"last_battle_time" binding:"required"`
	LogoutAt             timestamp `json:"logout_at" binding:"required" `
	Nickname             string    `json:"nickname" binding:"required"`
	UpdatedAt            timestamp `json:"updated_at" binding:"required"`
	structure.Statistics `json:"statistics" binding:"required"`
	*structure.Private   `json:"private" `
}

type Data map[string]PlayerDetail

type detailResponse struct {
	wgapi.ResponseBase
	Data Data `json:"data" `
}

func PersonalData(accountIds string) (results []PlayerDetail) {

	request, err := wgapi.Request(buildPath(personal_data), map[string]string{"account_id": accountIds})
	if err != nil {
		return nil
	}
	resp := &detailResponse{}
	json_err := json.Unmarshal(request, resp)
	if json_err != nil {
		fmt.Println(json_err)
		return nil
	}
	if resp.Status == "ok" {
		for _, val := range resp.Data {
			results = append(results, val)
		}
		return
	}

	return
}
