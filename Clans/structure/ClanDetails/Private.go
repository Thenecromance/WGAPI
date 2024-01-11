package structure
type Private struct {
OnlineMembers []int32/*list of integers*/ `json:"online_members" xml:"online_members" ` // List of clan members with an active game session in World of Tanks.

Treasury int32/*numeric*/ `json:"treasury" xml:"treasury" ` // Amount of gold in the сlan Treasury

ClanTreasury *ClanTreasury`json:"clan_treasury" xml:"clan_treasury" ` // Clan Treasury
}