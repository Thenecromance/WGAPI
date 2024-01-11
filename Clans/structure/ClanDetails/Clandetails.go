package structure

type Clandetails struct {
	AcceptsJoinRequests bool/*boolean*/ `json:"accepts_join_requests" xml:"accepts_join_requests" ` // Clan can invite players

	ClanId int32/*numeric*/ `json:"clan_id" xml:"clan_id" ` // Clan ID

	Color string/*string*/ `json:"color" xml:"color" ` // Clan color in HEX #RRGGBB

	CreatedAt int64/*timestamp*/ `json:"created_at" xml:"created_at" ` // Clan creation date

	CreatorId int32/*numeric*/ `json:"creator_id" xml:"creator_id" ` // Clan creator ID

	CreatorName string/*string*/ `json:"creator_name" xml:"creator_name" ` // Clan creator's name

	Description string/*string*/ `json:"description" xml:"description" ` // Clan description

	DescriptionHtml string/*string*/ `json:"description_html" xml:"description_html" ` // Clan description in HTML

	IsClanDisbanded bool/*boolean*/ `json:"is_clan_disbanded" xml:"is_clan_disbanded" ` // Clan has been deleted. The deleted clan data contains valid values for the following fields only:

	LeaderId int32/*numeric*/ `json:"leader_id" xml:"leader_id" ` // Clan Commander ID

	LeaderName string/*string*/ `json:"leader_name" xml:"leader_name" ` // Commander's name

	MembersCount int32/*numeric*/ `json:"members_count" xml:"members_count" ` // Number of clan members

	Motto string/*string*/ `json:"motto" xml:"motto" ` // Clan motto

	Name string/*string*/ `json:"name" xml:"name" ` // Clan name

	OldName string/*string*/ `json:"old_name" xml:"old_name" ` // Old clan name

	OldTag string/*string*/ `json:"old_tag" xml:"old_tag" ` // Old clan tag

	RenamedAt int64/*timestamp*/ `json:"renamed_at" xml:"renamed_at" ` // Time (UTC) when clan name was changed

	Tag string/*string*/ `json:"tag" xml:"tag" ` // Clan tag

	UpdatedAt int64/*timestamp*/ `json:"updated_at" xml:"updated_at" ` // Time when clan details were updated

	Emblems *Emblems `json:"emblems" xml:"emblems" ` // Information on clan emblems in games and on clan portal
	Members *Members `json:"members" xml:"members" ` // Information on clan members. Field format depends on members_key input parameter.
	Private *Private `json:"private" xml:"private" ` // Restricted clan information
}
