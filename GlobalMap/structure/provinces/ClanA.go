package structure

type ClanA struct {
	BattleReward int32 /*numeric*/ `json:"battle_reward" xml:"battle_reward" ` // Award

	ClanId int32 /*numeric*/ `json:"clan_id" xml:"clan_id" ` // Clan ID

	LooseEloDelta int32 /*numeric*/ `json:"loose_elo_delta" xml:"loose_elo_delta" ` // Changes in Elo-rating due to defeat

	WinEloDelta int32 /*numeric*/ `json:"win_elo_delta" xml:"win_elo_delta" ` // Changes in Elo-rating due to victory

}
