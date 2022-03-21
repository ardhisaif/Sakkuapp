package plannedpayment

type InputUser struct {
	Description string `json:"description"`
	CategoryID    string  `json:"category_id"`
	Price       float64 `json:"price"`
}
