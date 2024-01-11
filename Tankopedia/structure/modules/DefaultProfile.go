package structure

type DefaultProfile struct {
	Engine     *Engine     `json:"engine" xml:"engine" `         // Engine characteristics
	Gun        *Gun        `json:"gun" xml:"gun" `               // Gun characteristics
	Radio      *Radio      `json:"radio" xml:"radio" `           // Radio characteristics
	Suspension *Suspension `json:"suspension" xml:"suspension" ` // Suspension characteristics
	Turret     *Turret     `json:"turret" xml:"turret" `         // Turret characteristics
}
