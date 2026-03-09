package models

type CardSet struct {
	SetName    string  `json:"set_name"`
	SetCode    string  `json:"set_code"`
	NumOfCards int     `json:"num_of_cards"`
	TCGDate    *string `json:"tcg_date,omitempty"`
}
