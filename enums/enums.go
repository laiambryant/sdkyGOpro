package enums

// Language represents a non-English language supported by the YGOProDeck API.
// Pass a Language value to [query.Query.Language] to request card data in that
// language.
type Language string

const (
	LanguageFr Language = "fr" // French
	LanguageDe Language = "de" // German
	LanguageIt Language = "it" // Italian
	LanguagePt Language = "pt" // Portuguese
)

// Attribute represents a Yu-Gi-Oh! monster attribute as used in the
// YGOProDeck API filter parameters.
type Attribute string

const (
	AttributeDark   Attribute = "dark"
	AttributeLight  Attribute = "light"
	AttributeFire   Attribute = "fire"
	AttributeWater  Attribute = "water"
	AttributeWind   Attribute = "wind"
	AttributeEarth  Attribute = "earth"
	AttributeDivine Attribute = "divine"
)

// Format represents a Yu-Gi-Oh! game format recognised by the YGOProDeck API.
type Format string

const (
	FormatTCG        Format = "tcg"
	FormatGoat       Format = "goat"
	FormatOCGGoat    Format = "ocg goat"
	FormatSpeedDuel  Format = "speed duel"
	FormatMasterDuel Format = "master duel"
	FormatRushDuel   Format = "rush duel"
	FormatDuelLinks  Format = "duel links"
)

// SortOrder represents a field by which the API can sort card results.
type SortOrder string

const (
	SortATK   SortOrder = "atk"
	SortDEF   SortOrder = "def"
	SortName  SortOrder = "name"
	SortType  SortOrder = "type"
	SortLevel SortOrder = "level"
	SortID    SortOrder = "id"
	SortNew   SortOrder = "new"
)

// Banlist identifies a Yu-Gi-Oh! banlist used to filter cards by their ban
// status.
type Banlist string

const (
	BanlistTCG  Banlist = "tcg"
	BanlistOCG  Banlist = "ocg"
	BanlistGoat Banlist = "goat"
)
