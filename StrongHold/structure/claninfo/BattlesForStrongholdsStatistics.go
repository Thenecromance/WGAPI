package structure
type BattlesForStrongholdsStatistics struct {
LastTime10 int64/*timestamp*/ `json:"last_time_10" xml:"last_time_10" ` // End time of the last battle held on Tier X vehicles

Lose10 int32/*numeric*/ `json:"lose_10" xml:"lose_10" ` // Number of defeats on Tier X vehicles

Total10 int32/*numeric*/ `json:"total_10" xml:"total_10" ` // Total number of battles on Tier X vehicles

Total10In28d int32/*numeric*/ `json:"total_10_in_28d" xml:"total_10_in_28d" ` // Total number of battles on Tier X vehicles within the last 28 days

Win10 int32/*numeric*/ `json:"win_10" xml:"win_10" ` // Number of victories on Tier X vehicles

Win10In28d int32/*numeric*/ `json:"win_10_in_28d" xml:"win_10_in_28d" ` // Number of victories on Tier X vehicles within the last 28 days

}