package structure

type Clans struct {
	ClanId int32 /*numeric*/ `json:"clan_id" xml:"clan_id" ` // Clan ID

	Color string /*string*/ `json:"color" xml:"color" ` // Clan color in HEX #RRGGBB

	CreatedAt int64 /*timestamp*/ `json:"created_at" xml:"created_at" ` // Clan creation date

	MembersCount int32 /*numeric*/ `json:"members_count" xml:"members_count" ` // Number of clan members

	Name string /*string*/ `json:"name" xml:"name" ` // Clan name

	Tag string /*string*/ `json:"tag" xml:"tag" ` // Clan tag

	Emblems *Emblems `json:"emblems" xml:"emblems" ` // Information on clan emblems in games and on clan portal
}
