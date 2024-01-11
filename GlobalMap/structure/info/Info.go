package structure

type Info struct {
	LastTurn int32/*numeric*/ `json:"last_turn" xml:"last_turn" ` // Number of last calculated turn

	LastTurnCalculatedAt int64/*timestamp*/ `json:"last_turn_calculated_at" xml:"last_turn_calculated_at" ` // Calculation time of the last turn in UTC

	LastTurnCreatedAt int64/*timestamp*/ `json:"last_turn_created_at" xml:"last_turn_created_at" ` // Creation time of the last calculated turn in UTC

	State string/*string*/ `json:"state" xml:"state" ` // Map status: active, frozen, turn_calculation_in_progress

}
