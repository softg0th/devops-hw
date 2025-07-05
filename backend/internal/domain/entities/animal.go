package entities

type NewAnimal struct {
	Name    string `json:"Name"`
	Type    string `json:"Type"`
	Color   string `json:"Color"`
	StoreID int    `json:"StoreID"`
	Age     int    `json:"Age"`
	Price   int    `json:"Price"`
}

type UpdatedAnimal struct {
	ID      int    `json:"ID"`
	Name    string `json:"Name"`
	Type    string `json:"Type"`
	Color   string `json:"Color"`
	StoreID int    `json:"StoreID"`
	Age     int    `json:"Age"`
	Price   int    `json:"Price"`
}

type Animal struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Type         string  `json:"type"`
	Color        string  `json:"color"`
	StoreID      int     `json:"store_id"`
	StoreAddress string  `json:"store_address"`
	Age          int     `json:"age"`
	Price        float64 `json:"price"`
}
