package models

import (
	"encoding/json"
	"testing"
)

func TestCardUnmarshal(t *testing.T) {
	data := `{
		"id": 46986414,
		"name": "Dark Magician",
		"type": "Normal Monster",
		"frameType": "normal",
		"desc": "The ultimate wizard.",
		"atk": 2500,
		"def": 2100,
		"level": 7,
		"race": "Spellcaster",
		"attribute": "DARK",
		"card_images": [{"id": 46986414, "image_url": "https://images.ygoprodeck.com/images/cards/46986414.jpg", "image_url_small": "https://images.ygoprodeck.com/images/cards_small/46986414.jpg", "image_url_cropped": "https://images.ygoprodeck.com/images/cards_cropped/46986414.jpg"}],
		"card_prices": [{"cardmarket_price": "0.42", "tcgplayer_price": "0.48", "ebay_price": "0.99", "amazon_price": "0.95", "coolstuffinc_price": "0.99"}],
		"ygoprodeck_url": "https://ygoprodeck.com/card/dark-magician-1561"
	}`
	var c Card
	if err := json.Unmarshal([]byte(data), &c); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if c.ID != 46986414 {
		t.Fatalf("expected ID 46986414, got %d", c.ID)
	}
	if c.Name != "Dark Magician" {
		t.Fatalf("expected name Dark Magician, got %s", c.Name)
	}
	if c.ATK == nil || *c.ATK != 2500 {
		t.Fatalf("expected ATK 2500")
	}
	if c.DEF == nil || *c.DEF != 2100 {
		t.Fatalf("expected DEF 2100")
	}
	if c.Level == nil || *c.Level != 7 {
		t.Fatalf("expected Level 7")
	}
	if c.Attribute == nil || *c.Attribute != "DARK" {
		t.Fatalf("expected Attribute DARK")
	}
	if len(c.CardImages) != 1 {
		t.Fatalf("expected 1 card image, got %d", len(c.CardImages))
	}
	if len(c.CardPrices) != 1 {
		t.Fatalf("expected 1 card price, got %d", len(c.CardPrices))
	}
}

func TestCardWithLinkMarkers(t *testing.T) {
	data := `{
		"id": 1,
		"name": "Link Monster",
		"type": "Link Monster",
		"frameType": "link",
		"desc": "A link monster.",
		"atk": 2000,
		"race": "Cyberse",
		"linkval": 3,
		"linkmarkers": ["Top", "Bottom-Left", "Bottom-Right"],
		"card_images": [],
		"card_prices": [],
		"ygoprodeck_url": "https://ygoprodeck.com/card/link-monster"
	}`
	var c Card
	if err := json.Unmarshal([]byte(data), &c); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if c.LinkVal == nil || *c.LinkVal != 3 {
		t.Fatalf("expected LinkVal 3")
	}
	if len(c.LinkMarkers) != 3 {
		t.Fatalf("expected 3 link markers, got %d", len(c.LinkMarkers))
	}
}

func TestCardWithBanlistInfo(t *testing.T) {
	data := `{
		"id": 2,
		"name": "Banned Card",
		"type": "Spell Card",
		"frameType": "spell",
		"desc": "A spell.",
		"race": "Normal",
		"banlist_info": {"ban_tcg": "Banned", "ban_ocg": "Limited"},
		"card_images": [],
		"card_prices": [],
		"ygoprodeck_url": "https://ygoprodeck.com/card/banned"
	}`
	var c Card
	if err := json.Unmarshal([]byte(data), &c); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if c.BanlistInfo == nil {
		t.Fatalf("expected banlist info")
	}
	if c.BanlistInfo.BanTCG == nil || *c.BanlistInfo.BanTCG != "Banned" {
		t.Fatalf("expected BanTCG Banned")
	}
	if c.BanlistInfo.BanOCG == nil || *c.BanlistInfo.BanOCG != "Limited" {
		t.Fatalf("expected BanOCG Limited")
	}
	if c.BanlistInfo.BanGoat != nil {
		t.Fatalf("expected BanGoat nil")
	}
}

func TestCardWithCardSets(t *testing.T) {
	data := `{
		"id": 3,
		"name": "Set Card",
		"type": "Trap Card",
		"frameType": "trap",
		"desc": "A trap.",
		"race": "Normal",
		"card_sets": [{"set_name": "Set A", "set_code": "SA-001", "set_rarity": "Common", "set_rarity_code": "(C)", "set_price": "1.00"}],
		"card_images": [],
		"card_prices": [],
		"ygoprodeck_url": "https://ygoprodeck.com/card/set-card"
	}`
	var c Card
	if err := json.Unmarshal([]byte(data), &c); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(c.CardSets) != 1 {
		t.Fatalf("expected 1 card set, got %d", len(c.CardSets))
	}
	if c.CardSets[0].SetName != "Set A" || c.CardSets[0].SetRarityCode != "(C)" {
		t.Fatalf("unexpected card set: %#v", c.CardSets[0])
	}
}

func TestCardSetUnmarshal(t *testing.T) {
	data := `{"set_name": "Legend of Blue Eyes", "set_code": "LOB", "num_of_cards": 126, "tcg_date": "2002-03-08"}`
	var cs CardSet
	if err := json.Unmarshal([]byte(data), &cs); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cs.SetName != "Legend of Blue Eyes" {
		t.Fatalf("expected set name, got %s", cs.SetName)
	}
	if cs.NumOfCards != 126 {
		t.Fatalf("expected 126 cards, got %d", cs.NumOfCards)
	}
	if cs.TCGDate == nil || *cs.TCGDate != "2002-03-08" {
		t.Fatalf("expected tcg_date 2002-03-08")
	}
}

func TestCardSetNoDate(t *testing.T) {
	data := `{"set_name": "Test Set", "set_code": "TS", "num_of_cards": 10}`
	var cs CardSet
	if err := json.Unmarshal([]byte(data), &cs); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cs.TCGDate != nil {
		t.Fatalf("expected nil tcg_date")
	}
}

func TestCardSetInfoUnmarshal(t *testing.T) {
	data := `{"id": 1, "name": "Dark Magician", "set_name": "LOB", "set_code": "LOB-005", "set_rarity": "Ultra Rare", "set_price": "5.00"}`
	var csi CardSetInfo
	if err := json.Unmarshal([]byte(data), &csi); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if csi.ID != 1 || csi.Name != "Dark Magician" || csi.SetRarity != "Ultra Rare" {
		t.Fatalf("unexpected card set info: %#v", csi)
	}
}

func TestArchetypeUnmarshal(t *testing.T) {
	data := `{"archetype_name": "Blue-Eyes"}`
	var a Archetype
	if err := json.Unmarshal([]byte(data), &a); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if a.ArchetypeName != "Blue-Eyes" {
		t.Fatalf("expected Blue-Eyes, got %s", a.ArchetypeName)
	}
}

func TestDBVersionUnmarshal(t *testing.T) {
	data := `{"database_version": "ver1.0", "last_update": "2024-01-15 10:30:00"}`
	var v DBVersion
	if err := json.Unmarshal([]byte(data), &v); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if v.DatabaseVersion != "ver1.0" || v.LastUpdate != "2024-01-15 10:30:00" {
		t.Fatalf("unexpected db version: %#v", v)
	}
}

func TestMetaUnmarshal(t *testing.T) {
	data := `{"current_rows": 10, "total_rows": 100, "rows_remaining": 90, "total_pages": 10, "pages_remaining": 9, "next_page": "http://next", "next_page_offset": 10}`
	var m Meta
	if err := json.Unmarshal([]byte(data), &m); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if m.CurrentRows != 10 || m.TotalRows != 100 || m.RowsRemaining != 90 {
		t.Fatalf("unexpected meta: %#v", m)
	}
	if m.NextPage == nil || *m.NextPage != "http://next" {
		t.Fatalf("expected next_page")
	}
	if m.NextPageOffset == nil || *m.NextPageOffset != 10 {
		t.Fatalf("expected next_page_offset 10")
	}
}

func TestMetaNoOptional(t *testing.T) {
	data := `{"current_rows": 5, "total_rows": 5, "rows_remaining": 0, "total_pages": 1, "pages_remaining": 0}`
	var m Meta
	if err := json.Unmarshal([]byte(data), &m); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if m.NextPage != nil || m.NextPageOffset != nil {
		t.Fatalf("expected nil optional fields")
	}
}

func TestCardResponseUnmarshal(t *testing.T) {
	data := `{"data": [{"id": 1, "name": "Test", "type": "Normal Monster", "frameType": "normal", "desc": "desc", "race": "Warrior", "card_images": [], "card_prices": [], "ygoprodeck_url": "url"}]}`
	var cr CardResponse
	if err := json.Unmarshal([]byte(data), &cr); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(cr.Data) != 1 {
		t.Fatalf("expected 1 card, got %d", len(cr.Data))
	}
	if cr.Meta != nil {
		t.Fatalf("expected nil meta")
	}
}

func TestCardResponseWithMeta(t *testing.T) {
	data := `{"data": [{"id": 1, "name": "Test", "type": "Normal Monster", "frameType": "normal", "desc": "desc", "race": "Warrior", "card_images": [], "card_prices": [], "ygoprodeck_url": "url"}], "meta": {"current_rows": 1, "total_rows": 100, "rows_remaining": 99, "total_pages": 100, "pages_remaining": 99}}`
	var cr CardResponse
	if err := json.Unmarshal([]byte(data), &cr); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cr.Meta == nil || cr.Meta.TotalRows != 100 {
		t.Fatalf("expected meta with total_rows 100")
	}
}

func TestCardImageUnmarshal(t *testing.T) {
	data := `{"id": 123, "image_url": "https://img/123.jpg", "image_url_small": "https://img/123_small.jpg", "image_url_cropped": "https://img/123_cropped.jpg"}`
	var ci CardImage
	if err := json.Unmarshal([]byte(data), &ci); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ci.ID != 123 || ci.ImageURL != "https://img/123.jpg" {
		t.Fatalf("unexpected card image: %#v", ci)
	}
}

func TestCardPriceUnmarshal(t *testing.T) {
	data := `{"cardmarket_price": "0.42", "tcgplayer_price": "0.48", "ebay_price": "0.99", "amazon_price": "0.95", "coolstuffinc_price": "0.99"}`
	var cp CardPrice
	if err := json.Unmarshal([]byte(data), &cp); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cp.CardmarketPrice != "0.42" || cp.TCGPlayerPrice != "0.48" {
		t.Fatalf("unexpected card price: %#v", cp)
	}
}
