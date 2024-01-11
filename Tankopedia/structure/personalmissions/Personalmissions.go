package structure

type Personalmissions struct {
	CampaignId int32/*numeric*/ `json:"campaign_id" xml:"campaign_id" ` // Campaign ID

	Description string/*string*/ `json:"description" xml:"description" ` // Campaign description

	Name string/*string*/ `json:"name" xml:"name" ` // Campaign name

	Operations *Operations `json:"operations" xml:"operations" ` // Campaign operations
}
