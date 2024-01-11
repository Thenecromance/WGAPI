package structure

type Claninfo struct {
	ClanId int32/*numeric*/ `json:"clan_id" xml:"clan_id" ` // Clan ID

	Name string/*string*/ `json:"name" xml:"name" ` // Clan name

	Tag string/*string*/ `json:"tag" xml:"tag" ` // Clan tag

	Private    *Private    `json:"private" xml:"private" `       // Restricted clan information on the Global Map
	Ratings    *Ratings    `json:"ratings" xml:"ratings" `       // Clan rating on the Global Map
	Statistics *Statistics `json:"statistics" xml:"statistics" ` // Clan statistics on the Global Map
}
