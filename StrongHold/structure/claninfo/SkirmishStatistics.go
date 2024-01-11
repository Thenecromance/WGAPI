package structure
type SkirmishStatistics struct {
LastTime10 int64/*timestamp*/ `json:"last_time_10" xml:"last_time_10" ` // End time of the last battle held on Tier X vehicles

LastTime6 int64/*timestamp*/ `json:"last_time_6" xml:"last_time_6" ` // End time of the last battle held on Tier VI vehicles

LastTime8 int64/*timestamp*/ `json:"last_time_8" xml:"last_time_8" ` // End time of the last battle held on Tier VIII vehicles

Lose10 int32/*numeric*/ `json:"lose_10" xml:"lose_10" ` // Number of defeats on Tier X vehicles

Lose6 int32/*numeric*/ `json:"lose_6" xml:"lose_6" ` // Number of defeats on Tier VI vehicles

Lose8 int32/*numeric*/ `json:"lose_8" xml:"lose_8" ` // Number of defeats on Tier VIII vehicles

Total10 int32/*numeric*/ `json:"total_10" xml:"total_10" ` // Total number of battles on Tier X vehicles

Total10In28d int32/*numeric*/ `json:"total_10_in_28d" xml:"total_10_in_28d" ` // Total number of battles on Tier X vehicles within the last 28 days

Total6 int32/*numeric*/ `json:"total_6" xml:"total_6" ` // Total number of battles on Tier VI vehicles

Total6In28d int32/*numeric*/ `json:"total_6_in_28d" xml:"total_6_in_28d" ` // Total number of battles on Tier VI vehicles within the last 28 days

Total8 int32/*numeric*/ `json:"total_8" xml:"total_8" ` // Total number of battles on Tier VIII vehicles

Total8In28d int32/*numeric*/ `json:"total_8_in_28d" xml:"total_8_in_28d" ` // Total number of battles on Tier VIII vehicles within the last 28 days

Win10 int32/*numeric*/ `json:"win_10" xml:"win_10" ` // Number of victories on Tier X vehicles

Win10In28d int32/*numeric*/ `json:"win_10_in_28d" xml:"win_10_in_28d" ` // Number of victories on Tier X vehicles within the last 28 days

Win6 int32/*numeric*/ `json:"win_6" xml:"win_6" ` // Number of victories on Tier VI vehicles

Win6In28d int32/*numeric*/ `json:"win_6_in_28d" xml:"win_6_in_28d" ` // Number of victories on Tier VI vehicles within the last 28 days

Win8 int32/*numeric*/ `json:"win_8" xml:"win_8" ` // Number of victories on Tier VIII vehicles

Win8In28d int32/*numeric*/ `json:"win_8_in_28d" xml:"win_8_in_28d" ` // Number of victories on Tier VIII vehicles within the last 28 days

}