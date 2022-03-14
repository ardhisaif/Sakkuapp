package transaction

// import "MyApp/datastore/model"

type InputUser struct {
	Description string `json:"description"`
	CategoryID    string  `json:"category_id"`
	Type        int8  `json:"type"`
	Price       float64 `json:"price"`
}
