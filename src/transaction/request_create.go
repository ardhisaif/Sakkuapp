package transaction

type InputUser struct {
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Type        string  `json:"type"`
	Price       float64 `json:"price"`
}

