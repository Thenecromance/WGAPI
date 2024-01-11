package structure

type Operations struct {
	Description string/*string*/ `json:"description" xml:"description" ` // Operation description

	Image string/*string*/ `json:"image" xml:"image" ` // Link to an operation image

	MissionsInSet int32/*numeric*/ `json:"missions_in_set" xml:"missions_in_set" ` // Number of missions in the branch

	Name string/*string*/ `json:"name" xml:"name" ` // Operation name

	NextId int32/*numeric*/ `json:"next_id" xml:"next_id" ` // Next operation ID

	OperationId int32/*numeric*/ `json:"operation_id" xml:"operation_id" ` // Operation ID

	SetsCount int32/*numeric*/ `json:"sets_count" xml:"sets_count" ` // Number of mission branches of the operation

	SetsToNext int32/*numeric*/ `json:"sets_to_next" xml:"sets_to_next" ` // Number of branches until the next operation

	Missions *Missions `json:"missions" xml:"missions" ` // Operation missions
	Reward   *Reward   `json:"reward" xml:"reward" `     // Reward for operation
}
