package structure
type Playersclanhistory struct {
AccountId int32/*numeric*/ `json:"account_id" xml:"account_id" ` // Player account ID

ClanId int32/*numeric*/ `json:"clan_id" xml:"clan_id" ` // Clan ID

JoinedAt int64/*timestamp*/ `json:"joined_at" xml:"joined_at" ` // Date when player joined clan

LeftAt int64/*timestamp*/ `json:"left_at" xml:"left_at" ` // Date when player left clan

Role string/*string*/ `json:"role" xml:"role" ` // Last position in clan

}