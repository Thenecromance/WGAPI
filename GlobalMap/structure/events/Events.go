package structure

type Events struct {
	End string/*string*/ `json:"end" xml:"end" ` // Finishing time

	EventId string/*string*/ `json:"event_id" xml:"event_id" ` // Event ID

	EventName string/*string*/ `json:"event_name" xml:"event_name" ` // Event name

	Start string/*string*/ `json:"start" xml:"start" ` // Start time

	Status string/*string*/ `json:"status" xml:"status" ` // Status

	Fronts *Fronts `json:"fronts" xml:"fronts" ` // Fronts. Only event fronts are presented in response.
}
