package structure

type Clanreserves struct {
	BonusType string/*string*/ `json:"bonus_type" xml:"bonus_type" ` // Reserve efficiency type

	Disposable bool/*boolean*/ `json:"disposable" xml:"disposable" ` // Indicates if the Reserve is a One-Time-Effect Reserve

	Icon string/*string*/ `json:"icon" xml:"icon" ` // URL to icon

	Name string/*string*/ `json:"name" xml:"name" ` // Reserve name

	Type string/*string*/ `json:"type" xml:"type" ` // Reserve type

	InStock *InStock `json:"in_stock" xml:"in_stock" ` // Available clan Reserves and their status
}
