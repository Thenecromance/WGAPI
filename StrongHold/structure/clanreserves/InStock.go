package structure
type InStock struct {
ActionTime int32/*numeric*/ `json:"action_time" xml:"action_time" ` // Duration of clan Reserves of each level

ActivatedAt int32/*numeric*/ `json:"activated_at" xml:"activated_at" ` // Activation time of clan Reserves of each level

ActiveTill int32/*numeric*/ `json:"active_till" xml:"active_till" ` // Expiration time of clan Reserves of each level

Amount int32/*numeric*/ `json:"amount" xml:"amount" ` // Number of clan Reserves of each level

Level int32/*numeric*/ `json:"level" xml:"level" ` // Level of available clan Reserves

Status string/*string*/ `json:"status" xml:"status" ` // Status of clan Reserves of each level

XLevelOnly bool/*boolean*/ `json:"x_level_only" xml:"x_level_only" ` // Indicates if the Reserve is only for Tier X vehicles

BonusValues *BonusValues`json:"bonus_values" xml:"bonus_values" ` // Reserve efficiencies
}