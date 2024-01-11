package structure
type ClanTreasury struct {
Credits int32/*numeric*/ `json:"credits" xml:"credits" ` // Amount of credits in the сlan Treasury

Crystal int32/*numeric*/ `json:"crystal" xml:"crystal" ` // Number of bonds in the сlan Treasury

Gold int32/*numeric*/ `json:"gold" xml:"gold" ` // Amount of gold in the сlan Treasury

}