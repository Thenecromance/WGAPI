package structure
type Info struct {
AccountId int32/*numeric*/ `json:"account_id" xml:"account_id" ` // Player account ID

ClanId int32/*numeric*/ `json:"clan_id" xml:"clan_id" ` // Clan ID

ClientLanguage string/*string*/ `json:"client_language" xml:"client_language" ` // Language selected in the game client

CreatedAt int64/*timestamp*/ `json:"created_at" xml:"created_at" ` // Date when player's account was created

GlobalRating int32/*numeric*/ `json:"global_rating" xml:"global_rating" ` // Personal rating

LastBattleTime int64/*timestamp*/ `json:"last_battle_time" xml:"last_battle_time" ` // Last battle time

LogoutAt int64/*timestamp*/ `json:"logout_at" xml:"logout_at" ` // End time of last game session

Nickname string/*string*/ `json:"nickname" xml:"nickname" ` // Player name

UpdatedAt int64/*timestamp*/ `json:"updated_at" xml:"updated_at" ` // Date when player details were updated

Private *Private`json:"private" xml:"private" ` // Player's private data
Statistics *Statistics`json:"statistics" xml:"statistics" ` // Player statistics
}

