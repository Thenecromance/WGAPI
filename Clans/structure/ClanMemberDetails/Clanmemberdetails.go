package structure
type Clanmemberdetails struct {
AccountId int32/*numeric*/ `json:"account_id" xml:"account_id" ` // Player account ID

AccountName string/*string*/ `json:"account_name" xml:"account_name" ` // Player name

JoinedAt int64/*timestamp*/ `json:"joined_at" xml:"joined_at" ` // Date when player joined clan

Role string/*string*/ `json:"role" xml:"role" ` // Technical position name

RoleI18n string/*string*/ `json:"role_i18n" xml:"role_i18n" ` // Position

Clan *Clan`json:"clan" xml:"clan" ` // Short clan info
}