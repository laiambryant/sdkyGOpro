package models

// CardSet represents a summary entry for a single Yu-Gi-Oh! card set as
// returned by /cardsets.php.
type CardSet struct {
	// SetName is the full name of the set.
	SetName string `json:"set_name"`
	// SetCode is the short code identifier (e.g. "LOB").
	SetCode string `json:"set_code"`
	// NumOfCards is the total number of cards in the set.
	NumOfCards int `json:"num_of_cards"`
	// TCGDate is the TCG release date in YYYY-MM-DD format. Nil if the set
	// has no known TCG release date.
	TCGDate *string `json:"tcg_date,omitempty"`
}
