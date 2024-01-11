package structure

type Stun struct {
	Duration []float32 /*object*/ `json:"duration" xml:"duration" ` // Stun duration (s) caused by this shell type, a list of values: min, max

}
