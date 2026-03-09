package enums

type Language string

const (
	LanguageFr Language = "fr"
	LanguageDe Language = "de"
	LanguageIt Language = "it"
	LanguagePt Language = "pt"
)

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

type Format string

const (
	FormatTCG       Format = "tcg"
	FormatGoat      Format = "goat"
	FormatOCGGoat   Format = "ocg goat"
	FormatSpeedDuel Format = "speed duel"
	FormatMasterDuel Format = "master duel"
	FormatRushDuel  Format = "rush duel"
	FormatDuelLinks Format = "duel links"
)

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

type Banlist string

const (
	BanlistTCG  Banlist = "tcg"
	BanlistOCG  Banlist = "ocg"
	BanlistGoat Banlist = "goat"
)
