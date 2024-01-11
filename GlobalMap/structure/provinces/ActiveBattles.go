package structure

type ActiveBattles struct {
	BattleReward int32/*numeric*/ `json:"battle_reward" xml:"battle_reward" ` // Award

	Round int32/*numeric*/ `json:"round" xml:"round" ` // Round

	StartAt string/*string*/ `json:"start_at" xml:"start_at" ` // Battle start time in UTC

	ClanA *ClanA `json:"clan_a" xml:"clan_a" ` // First challenging clan ID
	ClanB *ClanB `json:"clan_b" xml:"clan_b" ` // Second challenging clan ID
}
