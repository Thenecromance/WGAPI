package structure
type GroupedContacts struct {
Blocked []int32/*list of integers*/ `json:"blocked" xml:"blocked" ` // The contact was added to the blacklist. Note that the blocked contact still belongs to contact groups or to the ungrouped list of contacts.

Groups map[string]string/*associative array*/ `json:"groups" xml:"groups" ` // Groups

Ignored []int32/*list of integers*/ `json:"ignored" xml:"ignored" ` // Ignored

Muted []int32/*list of integers*/ `json:"muted" xml:"muted" ` // Muted

Ungrouped []int32/*list of integers*/ `json:"ungrouped" xml:"ungrouped" ` // Not grouped

}