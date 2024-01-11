package structure
type Images struct {
BigIcon string/*string*/ `json:"big_icon" xml:"big_icon" ` // URL to 160 x 100 px image of vehicle

ContourIcon string/*string*/ `json:"contour_icon" xml:"contour_icon" ` // URL to outline image of vehicle

SmallIcon string/*string*/ `json:"small_icon" xml:"small_icon" ` // URL to 124 x 31 px image of vehicle

}