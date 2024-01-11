package structure
type Achievements struct {
Achievements map[string]string/*associative array*/ `json:"achievements" xml:"achievements" ` // Achievements earned

Frags map[string]string/*associative array*/ `json:"frags" xml:"frags" ` // Achievement progress

MaxSeries map[string]string/*associative array*/ `json:"max_series" xml:"max_series" ` // Maximum values of achievement series

}