package structure
type Company struct {
BattleAvgXp int32/*numeric*/ `json:"battle_avg_xp" xml:"battle_avg_xp" ` // Average experience per battle

Battles int32/*numeric*/ `json:"battles" xml:"battles" ` // Battles fought

BattlesOnStunningVehicles int32/*numeric*/ `json:"battles_on_stunning_vehicles" xml:"battles_on_stunning_vehicles" ` // Number of battles on vehicles that cause the stun effect

CapturePoints int32/*numeric*/ `json:"capture_points" xml:"capture_points" ` // Base capture points

DamageDealt int32/*numeric*/ `json:"damage_dealt" xml:"damage_dealt" ` // Damage caused

DamageReceived int32/*numeric*/ `json:"damage_received" xml:"damage_received" ` // Damage received

Draws int32/*numeric*/ `json:"draws" xml:"draws" ` // Draws

DroppedCapturePoints int32/*numeric*/ `json:"dropped_capture_points" xml:"dropped_capture_points" ` // Base defense points

Frags int32/*numeric*/ `json:"frags" xml:"frags" ` // Vehicles destroyed

Hits int32/*numeric*/ `json:"hits" xml:"hits" ` // Hits

HitsPercents int32/*numeric*/ `json:"hits_percents" xml:"hits_percents" ` // Hit ratio

Losses int32/*numeric*/ `json:"losses" xml:"losses" ` // Defeats

Shots int32/*numeric*/ `json:"shots" xml:"shots" ` // Shots fired

Spotted int32/*numeric*/ `json:"spotted" xml:"spotted" ` // Enemies spotted

StunAssistedDamage int32/*numeric*/ `json:"stun_assisted_damage" xml:"stun_assisted_damage" ` // Damage to enemy vehicles stunned by you

StunNumber int32/*numeric*/ `json:"stun_number" xml:"stun_number" ` // Number of times an enemy was stunned by you

SurvivedBattles int32/*numeric*/ `json:"survived_battles" xml:"survived_battles" ` // Battles survived

Wins int32/*numeric*/ `json:"wins" xml:"wins" ` // Victories

Xp int32/*numeric*/ `json:"xp" xml:"xp" ` // Total experience

}