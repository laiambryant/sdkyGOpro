package models

type CardSetInfo struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	SetName   string `json:"set_name"`
	SetCode   string `json:"set_code"`
	SetRarity string `json:"set_rarity"`
	SetPrice  string `json:"set_price"`
}
