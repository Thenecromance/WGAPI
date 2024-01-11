package structure

type Ratings struct {
	Elo10 int32/*numeric*/ `json:"elo_10" xml:"elo_10" ` // Clan Elo rating in Absolute division

	Elo6 int32/*numeric*/ `json:"elo_6" xml:"elo_6" ` // Clan Elo rating in Medium division

	Elo8 int32/*numeric*/ `json:"elo_8" xml:"elo_8" ` // Clan Elo rating in Champion division

	UpdatedAt int64/*timestamp*/ `json:"updated_at" xml:"updated_at" ` // Ratings update time

}
