package structure
type GlobalRatingWeightedAvg struct {
Rank int32/*numeric*/ `json:"rank" xml:"rank" ` // Position

RankDelta int32/*numeric*/ `json:"rank_delta" xml:"rank_delta" ` // Change of position in rating

Value float64/*float*/ `json:"value" xml:"value" ` // Absolute value

}