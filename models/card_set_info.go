package models

// CardSetInfo represents a single card's membership in a set, as returned by
// /cardsetsinfo.php. Each element describes one printing of one card within the
// queried set.
type CardSetInfo struct {
	// ID is the YGOProDeck card ID.
	ID int `json:"id"`
	// Name is the card name.
	Name string `json:"name"`
	// SetName is the name of the set.
	SetName string `json:"set_name"`
	// SetCode is the card's set-specific code (e.g. "LOB-005").
	SetCode string `json:"set_code"`
	// SetRarity is the rarity label (e.g. "Ultra Rare").
	SetRarity string `json:"set_rarity"`
	// SetPrice is the current market price string for this printing.
	SetPrice string `json:"set_price"`
}
