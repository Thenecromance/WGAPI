package structure

type Boosters struct {
	BoosterId int32/*numeric*/ `json:"booster_id" xml:"booster_id" ` // Personal Reserve ID

	Description string/*string*/ `json:"description" xml:"description" ` // Personal Reserve description

	ExpiresAt int64/*timestamp*/ `json:"expires_at" xml:"expires_at" ` // Personal Reserve expiration time in UTC

	IsAuto bool/*boolean*/ `json:"is_auto" xml:"is_auto" ` // Personal Reserve auto activation flag

	Lifetime int32/*numeric*/ `json:"lifetime" xml:"lifetime" ` // Personal Reserve duration in seconds

	Name string/*string*/ `json:"name" xml:"name" ` // Personal Reserve name

	PriceCredit int32/*numeric*/ `json:"price_credit" xml:"price_credit" ` // Cost in credits

	PriceGold int32/*numeric*/ `json:"price_gold" xml:"price_gold" ` // Price in gold

	Resource string/*string*/ `json:"resource" xml:"resource" ` // Resource affected by Personal Reserve. Valid values: credits, experience, crew_experience, free_experience.

	Images *Images `json:"images" xml:"images" ` // Personal Reserve image
}
