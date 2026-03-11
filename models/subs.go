package models

// CardImage holds the URLs for a single print image of a card.
type CardImage struct {
	// ID is the image identifier, typically matching the card ID.
	ID int `json:"id"`
	// ImageURL is the URL of the full-size card image.
	ImageURL string `json:"image_url"`
	// ImageURLSmall is the URL of the small card image.
	ImageURLSmall string `json:"image_url_small"`
	// ImageURLCropped is the URL of the cropped card art image.
	ImageURLCropped string `json:"image_url_cropped"`
}

// CardPrice holds the latest market prices for a card from multiple vendors.
// All price values are decimal strings (e.g. "0.42").
type CardPrice struct {
	CardmarketPrice   string `json:"cardmarket_price"`
	TCGPlayerPrice    string `json:"tcgplayer_price"`
	EbayPrice         string `json:"ebay_price"`
	AmazonPrice       string `json:"amazon_price"`
	CoolStuffIncPrice string `json:"coolstuffinc_price"`
}

// CardSetEntry describes a single printing of a card within a set, embedded
// in a [Card].
type CardSetEntry struct {
	// SetName is the name of the set.
	SetName string `json:"set_name"`
	// SetCode is the card's set-specific code (e.g. "LOB-005").
	SetCode string `json:"set_code"`
	// SetRarity is the rarity label (e.g. "Ultra Rare").
	SetRarity string `json:"set_rarity"`
	// SetRarityCode is the short rarity code (e.g. "(UR)").
	SetRarityCode string `json:"set_rarity_code"`
	// SetPrice is the current market price string for this printing.
	SetPrice string `json:"set_price"`
}

// BanlistInfo describes a card's status on the TCG, OCG, and GOAT banlists.
// A nil pointer field means the card has no restriction in that list.
type BanlistInfo struct {
	// BanTCG is the TCG ban status (e.g. "Banned", "Limited", "Semi-Limited").
	BanTCG *string `json:"ban_tcg,omitempty"`
	// BanOCG is the OCG ban status.
	BanOCG *string `json:"ban_ocg,omitempty"`
	// BanGoat is the GOAT format ban status.
	BanGoat *string `json:"ban_goat,omitempty"`
}

// Meta contains pagination metadata returned alongside card results when the
// num or offset parameters are used.
type Meta struct {
	// CurrentRows is the number of cards in the current page.
	CurrentRows int `json:"current_rows"`
	// TotalRows is the total number of cards matching the query.
	TotalRows int `json:"total_rows"`
	// RowsRemaining is the number of cards not yet returned.
	RowsRemaining int `json:"rows_remaining"`
	// TotalPages is the total number of pages.
	TotalPages int `json:"total_pages"`
	// PagesRemaining is the number of pages not yet returned.
	PagesRemaining int `json:"pages_remaining"`
	// NextPage is the URL of the next page, if one exists.
	NextPage *string `json:"next_page,omitempty"`
	// NextPageOffset is the offset value to use to fetch the next page.
	NextPageOffset *int `json:"next_page_offset,omitempty"`
}

// CardResponse is the raw envelope returned by /cardinfo.php and
// /randomcard.php. SDK methods unwrap this into [CardInfoResponse].
type CardResponse struct {
	Data []Card `json:"data"`
	Meta *Meta  `json:"meta,omitempty"`
}

// CardInfoResponse is the value returned by [ygoprodeck.YGOProDeck.GetCards].
// Meta is nil when the API response did not include pagination metadata.
type CardInfoResponse struct {
	// Cards holds the matched cards.
	Cards []Card
	// Meta holds optional pagination metadata.
	Meta *Meta
}
