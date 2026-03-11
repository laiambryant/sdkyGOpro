package models

// Card represents a single Yu-Gi-Oh! card as returned by /cardinfo.php and
// /randomcard.php. Fields that are not present for every card type (e.g. ATK
// for Spell/Trap cards, or LinkVal for non-Link monsters) are pointer types so
// callers can distinguish an absent value from a zero value.
type Card struct {
	// ID is the YGOProDeck card ID.
	ID int `json:"id"`
	// Name is the card name.
	Name string `json:"name"`
	// Type is the card type string (e.g. "Normal Monster", "Spell Card").
	Type string `json:"type"`
	// FrameType is the card frame category (e.g. "normal", "effect", "spell").
	FrameType string `json:"frameType"`
	// Desc is the card text / effect description.
	Desc string `json:"desc"`
	// ATK is the monster's attack value. Nil for non-monster cards.
	ATK *int `json:"atk,omitempty"`
	// DEF is the monster's defence value. Nil for Link monsters and non-monsters.
	DEF *int `json:"def,omitempty"`
	// Level is the monster level or rank. Nil for non-monsters and Link monsters.
	Level *int `json:"level,omitempty"`
	// Race is the monster sub-type (e.g. "Spellcaster") or spell/trap type
	// (e.g. "Normal", "Continuous").
	Race string `json:"race"`
	// Attribute is the monster attribute (e.g. "DARK"). Nil for non-monsters.
	Attribute *string `json:"attribute,omitempty"`
	// Scale is the Pendulum scale value. Nil for non-Pendulum cards.
	Scale *int `json:"scale,omitempty"`
	// LinkVal is the Link monster rating. Nil for non-Link cards.
	LinkVal *int `json:"linkval,omitempty"`
	// LinkMarkers lists the Link arrow directions for Link monsters.
	LinkMarkers []string `json:"linkmarkers,omitempty"`
	// CardSets lists the sets and rarities in which this card appears.
	CardSets []CardSetEntry `json:"card_sets,omitempty"`
	// CardImages contains all available print images for this card.
	CardImages []CardImage `json:"card_images"`
	// CardPrices contains the latest market prices from multiple vendors.
	CardPrices []CardPrice `json:"card_prices"`
	// BanlistInfo contains the card's ban status across formats. Nil if
	// the card is not on any banlist.
	BanlistInfo *BanlistInfo `json:"banlist_info,omitempty"`
	// YGOProDeckURL is the canonical URL for this card on ygoprodeck.com.
	YGOProDeckURL string `json:"ygoprodeck_url"`
}
