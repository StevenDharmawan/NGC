package models

type Inventories struct {
	ItemCode    int64  `json:"item_code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Stock       int    `json:"stock"`
	Status      string `json:"status"`
}
