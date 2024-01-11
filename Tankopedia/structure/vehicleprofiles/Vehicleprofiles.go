package structure
type Vehicleprofiles struct {
IsDefault bool/*boolean*/ `json:"is_default" xml:"is_default" ` // Standard configuration

PriceCredit int32/*numeric*/ `json:"price_credit" xml:"price_credit" ` // Cost in credits

ProfileId string/*string*/ `json:"profile_id" xml:"profile_id" ` // Vehicle Configuration ID

TankId int32/*numeric*/ `json:"tank_id" xml:"tank_id" ` // Vehicle ID

}