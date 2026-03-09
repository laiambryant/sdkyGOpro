package models

type Card struct {
	ID            int            `json:"id"`
	Name          string         `json:"name"`
	Type          string         `json:"type"`
	FrameType     string         `json:"frameType"`
	Desc          string         `json:"desc"`
	ATK           *int           `json:"atk,omitempty"`
	DEF           *int           `json:"def,omitempty"`
	Level         *int           `json:"level,omitempty"`
	Race          string         `json:"race"`
	Attribute     *string        `json:"attribute,omitempty"`
	Scale         *int           `json:"scale,omitempty"`
	LinkVal       *int           `json:"linkval,omitempty"`
	LinkMarkers   []string       `json:"linkmarkers,omitempty"`
	CardSets      []CardSetEntry `json:"card_sets,omitempty"`
	CardImages    []CardImage    `json:"card_images"`
	CardPrices    []CardPrice    `json:"card_prices"`
	BanlistInfo   *BanlistInfo   `json:"banlist_info,omitempty"`
	YGOProDeckURL string         `json:"ygoprodeck_url"`
}
