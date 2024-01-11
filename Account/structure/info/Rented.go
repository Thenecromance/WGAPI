package structure
type Rented struct {
CompensationCredits int32/*numeric*/ `json:"compensation_credits" xml:"compensation_credits" ` // Rental compensation in credits

CompensationGold int32/*numeric*/ `json:"compensation_gold" xml:"compensation_gold" ` // Rental compensation in gold

ExpirationTime int64/*timestamp*/ `json:"expiration_time" xml:"expiration_time" ` // Vehicle Rental expiration time

TankId int32/*numeric*/ `json:"tank_id" xml:"tank_id" ` // Rented vehicle ID

}