package structure
type Modules struct {
EngineId int32/*numeric*/ `json:"engine_id" xml:"engine_id" ` // Engine ID

GunId int32/*numeric*/ `json:"gun_id" xml:"gun_id" ` // Gun ID

RadioId int32/*numeric*/ `json:"radio_id" xml:"radio_id" ` // Radio ID

SuspensionId int32/*numeric*/ `json:"suspension_id" xml:"suspension_id" ` // Suspension ID

TurretId int32/*numeric*/ `json:"turret_id" xml:"turret_id" ` // Turret ID

}