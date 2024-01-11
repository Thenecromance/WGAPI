package structure
type Private struct {
BanInfo string/*string*/ `json:"ban_info" xml:"ban_info" ` // Account ban details

BanTime int64/*timestamp*/ `json:"ban_time" xml:"ban_time" ` // End time of account ban

BattleLifeTime int32/*numeric*/ `json:"battle_life_time" xml:"battle_life_time" ` // Overall battle life time in seconds

Bonds int32/*numeric*/ `json:"bonds" xml:"bonds" ` // Bonds

Credits int32/*numeric*/ `json:"credits" xml:"credits" ` // Credits

FreeXp int32/*numeric*/ `json:"free_xp" xml:"free_xp" ` // Free Experience

Garage []int32/*list of integers*/ `json:"garage" xml:"garage" ` // Vehicles in the Garage.

Gold int32/*numeric*/ `json:"gold" xml:"gold" ` // Gold

IsBoundToPhone bool/*boolean*/ `json:"is_bound_to_phone" xml:"is_bound_to_phone" ` // Indicates if mobile phone number was added to the account

IsPremium bool/*boolean*/ `json:"is_premium" xml:"is_premium" ` // Indicates if the account is Premium Account

PersonalMissions map[string]string/*associative array*/ `json:"personal_missions" xml:"personal_missions" ` // Personal Missions Progress. The key is a task id, the value is a status.

PremiumExpiresAt int64/*timestamp*/ `json:"premium_expires_at" xml:"premium_expires_at" ` // Premium Account expiration time

Boosters *Boosters`json:"boosters" xml:"boosters" ` // Personal Reserves.
GroupedContacts *GroupedContacts`json:"grouped_contacts" xml:"grouped_contacts" ` // Contact groups.
Rented *Rented`json:"rented" xml:"rented" ` // Vehicle Rental.
Restrictions *Restrictions`json:"restrictions" xml:"restrictions" ` // Account restrictions
}