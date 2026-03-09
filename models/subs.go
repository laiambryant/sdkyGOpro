package models

type CardImage struct {
	ID              int    `json:"id"`
	ImageURL        string `json:"image_url"`
	ImageURLSmall   string `json:"image_url_small"`
	ImageURLCropped string `json:"image_url_cropped"`
}

type CardPrice struct {
	CardmarketPrice   string `json:"cardmarket_price"`
	TCGPlayerPrice    string `json:"tcgplayer_price"`
	EbayPrice         string `json:"ebay_price"`
	AmazonPrice       string `json:"amazon_price"`
	CoolStuffIncPrice string `json:"coolstuffinc_price"`
}

type CardSetEntry struct {
	SetName       string `json:"set_name"`
	SetCode       string `json:"set_code"`
	SetRarity     string `json:"set_rarity"`
	SetRarityCode string `json:"set_rarity_code"`
	SetPrice      string `json:"set_price"`
}

type BanlistInfo struct {
	BanTCG  *string `json:"ban_tcg,omitempty"`
	BanOCG  *string `json:"ban_ocg,omitempty"`
	BanGoat *string `json:"ban_goat,omitempty"`
}

type Meta struct {
	CurrentRows    int     `json:"current_rows"`
	TotalRows      int     `json:"total_rows"`
	RowsRemaining  int     `json:"rows_remaining"`
	TotalPages     int     `json:"total_pages"`
	PagesRemaining int     `json:"pages_remaining"`
	NextPage       *string `json:"next_page,omitempty"`
	NextPageOffset *int    `json:"next_page_offset,omitempty"`
}

type CardResponse struct {
	Data []Card `json:"data"`
	Meta *Meta  `json:"meta,omitempty"`
}

type CardInfoResponse struct {
	Cards []Card
	Meta  *Meta
}
