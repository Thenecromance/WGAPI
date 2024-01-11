package structure
type Radio struct {
Name string/*string*/ `json:"name" xml:"name" ` // Module name

SignalRange int32/*numeric*/ `json:"signal_range" xml:"signal_range" ` // Signal range

Tag string/*string*/ `json:"tag" xml:"tag" ` // Module tag

Tier int32/*numeric*/ `json:"tier" xml:"tier" ` // Tier

Weight int32/*numeric*/ `json:"weight" xml:"weight" ` // Weight (kg)

}