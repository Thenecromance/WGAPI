package structure

type Seasons struct {
	End string/*string*/ `json:"end" xml:"end" ` // Finishing time

	SeasonId string/*string*/ `json:"season_id" xml:"season_id" ` // Season ID

	SeasonName string/*string*/ `json:"season_name" xml:"season_name" ` // Season name

	Start string/*string*/ `json:"start" xml:"start" ` // Start time

	Status string/*string*/ `json:"status" xml:"status" ` // Status

	Fronts *Fronts `json:"fronts" xml:"fronts" ` // Fronts. Only season fronts are presented in response.
}
