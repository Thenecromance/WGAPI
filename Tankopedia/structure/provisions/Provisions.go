package structure

type Provisions struct {
	Description string/*string*/ `json:"description" xml:"description" ` // Achievement description

	Image string/*string*/ `json:"image" xml:"image" ` // Image link

	Name string/*string*/ `json:"name" xml:"name" ` // Vehicle name

	PriceCredit int32/*numeric*/ `json:"price_credit" xml:"price_credit" ` // Cost in credits

	PriceGold int32/*numeric*/ `json:"price_gold" xml:"price_gold" ` // Cost in gold

	ProvisionId int32/*numeric*/ `json:"provision_id" xml:"provision_id" ` // Equipment or consumables ID

	Tag string/*string*/ `json:"tag" xml:"tag" ` // Technical name

	Type string/*string*/ `json:"type" xml:"type" ` // Type: consumable or equipment

	Weight int32/*numeric*/ `json:"weight" xml:"weight" ` // Weight in kilos. Used for equipment only.

}
