package structure
type Rewards struct {
Berths int32/*numeric*/ `json:"berths" xml:"berths" ` // Bunks in Barracks

Conditions string/*string*/ `json:"conditions" xml:"conditions" ` // Mission conditions

Credits int32/*numeric*/ `json:"credits" xml:"credits" ` // Credits

FreeXp int32/*numeric*/ `json:"free_xp" xml:"free_xp" ` // Free Experience

Items map[string]string/*associative array*/ `json:"items" xml:"items" ` // List of equipment or consumables in the following format: Id–number of items

Premium int32/*numeric*/ `json:"premium" xml:"premium" ` // Days of Premium Account

Slots int32/*numeric*/ `json:"slots" xml:"slots" ` // Slots

Tokens int32/*numeric*/ `json:"tokens" xml:"tokens" ` // Tokens

}