package structure
type Crew struct {
MemberId string/*string*/ `json:"member_id" xml:"member_id" ` // Crew member ID

Roles map[string]string/*associative array*/ `json:"roles" xml:"roles" ` // List of crew member roles

}