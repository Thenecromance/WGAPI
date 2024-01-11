package structure
type Vehicles struct {
Description string/*string*/ `json:"description" xml:"description" ` // Vehicle description

Engines []int32/*list of integers*/ `json:"engines" xml:"engines" ` // List of compatible engine IDs

Guns []int32/*list of integers*/ `json:"guns" xml:"guns" ` // List of compatible gun IDs

IsGift bool/*boolean*/ `json:"is_gift" xml:"is_gift" ` // Indicates if the vehicle is a gift vehicle

IsPremium bool/*boolean*/ `json:"is_premium" xml:"is_premium" ` // Indicates if the vehicle is Premium vehicle

IsPremiumIgr bool/*boolean*/ `json:"is_premium_igr" xml:"is_premium_igr" ` // Indicates the IGR vehicle. Active only for Korea realm

IsWheeled bool/*boolean*/ `json:"is_wheeled" xml:"is_wheeled" ` // Indicates if the vehicle is a wheeled vehicle

Name string/*string*/ `json:"name" xml:"name" ` // Vehicle name

Nation string/*string*/ `json:"nation" xml:"nation" ` // Nation

NextTanks map[string]string/*associative array*/ `json:"next_tanks" xml:"next_tanks" ` // List of vehicles available for research in form of pairs:

PriceCredit int32/*numeric*/ `json:"price_credit" xml:"price_credit" ` // Cost in credits

PriceGold int32/*numeric*/ `json:"price_gold" xml:"price_gold" ` // Cost in gold

PricesXp map[string]string/*associative array*/ `json:"prices_xp" xml:"prices_xp" ` // List of research costs in form of pairs:

Provisions []int32/*list of integers*/ `json:"provisions" xml:"provisions" ` // List of IDs of compatible equipment and consumables

Radios []int32/*list of integers*/ `json:"radios" xml:"radios" ` // List of compatible radio IDs

ShortName string/*string*/ `json:"short_name" xml:"short_name" ` // Vehicle short name

Suspensions []int32/*list of integers*/ `json:"suspensions" xml:"suspensions" ` // List of compatible suspension IDs

Tag string/*string*/ `json:"tag" xml:"tag" ` // Vehicle tag

TankId int32/*numeric*/ `json:"tank_id" xml:"tank_id" ` // Vehicle ID

Tier int32/*numeric*/ `json:"tier" xml:"tier" ` // Tier

Turrets []int32/*list of integers*/ `json:"turrets" xml:"turrets" ` // List of compatible turret IDs

Type string/*string*/ `json:"type" xml:"type" ` // Vehicle type

Crew *Crew`json:"crew" xml:"crew" ` // Crew
DefaultProfile *DefaultProfile`json:"default_profile" xml:"default_profile" ` // Standard configuration characteristics
Images *Images`json:"images" xml:"images" ` // Image links
ModulesTree *ModulesTree`json:"modules_tree" xml:"modules_tree" ` // Module research information
Multination *Multination`json:"multination" xml:"multination" ` // Информация об мультинации
}