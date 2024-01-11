package structure

type Clanbattles struct {
	AttackType string/*string*/ `json:"attack_type" xml:"attack_type" ` // Attack Type: ground, auction, tournament

	CompetitorId int32/*numeric*/ `json:"competitor_id" xml:"competitor_id" ` // Enemy clan ID

	FrontId string/*string*/ `json:"front_id" xml:"front_id" ` // Front ID

	FrontName string/*string*/ `json:"front_name" xml:"front_name" ` // Front name

	ProvinceId string/*string*/ `json:"province_id" xml:"province_id" ` // Province ID

	ProvinceName string/*string*/ `json:"province_name" xml:"province_name" ` // Province name

	Time int64/*timestamp*/ `json:"time" xml:"time" ` // Battle date and time

	Type string/*string*/ `json:"type" xml:"type" ` // Battle type: attack or defense

	VehicleLevel int32/*numeric*/ `json:"vehicle_level" xml:"vehicle_level" ` // Vehicle Tier

}
