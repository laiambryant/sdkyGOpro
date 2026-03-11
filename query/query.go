package query

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/laiambryant/sdkyGOprodeck/enums"
)

// Query is a fluent builder for /cardinfo.php query parameters. Create one
// with [New], chain filter methods, then pass it to
// [ygoprodeck.YGOProDeck.GetCards]. All methods return the same *Query pointer
// so calls can be chained.
type Query struct {
	params []param
}

type param struct {
	key   string
	value string
}

// New returns an empty Query.
func New() *Query {
	return &Query{}
}

func (q *Query) add(key, value string) {
	q.params = append(q.params, param{key, value})
}

// Name filters cards whose name matches name exactly (case-insensitive).
func (q *Query) Name(name string) *Query {
	q.add("name", name)
	return q
}

// FuzzyName filters cards whose name contains fname (partial match).
func (q *Query) FuzzyName(fname string) *Query {
	q.add("fname", fname)
	return q
}

// ID filters by the YGOProDeck card ID.
func (q *Query) ID(id int) *Query {
	q.add("id", fmt.Sprintf("%d", id))
	return q
}

// KonamiID filters by the official Konami card ID.
func (q *Query) KonamiID(id int) *Query {
	q.add("konami_id", fmt.Sprintf("%d", id))
	return q
}

// CardType filters by card type (e.g. "Normal Monster", "Spell Card").
func (q *Query) CardType(ct string) *Query {
	q.add("type", ct)
	return q
}

// Race filters by the card's race / sub-type (e.g. "Spellcaster", "Dragon").
func (q *Query) Race(race string) *Query {
	q.add("race", race)
	return q
}

// Attribute filters by monster attribute.
func (q *Query) Attribute(attr enums.Attribute) *Query {
	q.add("attribute", string(attr))
	return q
}

// ATK filters by attack value. The API supports operator prefixes such as
// "gte2500" or "lte1000" as well as exact numeric strings.
func (q *Query) ATK(val string) *Query {
	q.add("atk", val)
	return q
}

// DEF filters by defence value. Supports the same operator prefixes as [ATK].
func (q *Query) DEF(val string) *Query {
	q.add("def", val)
	return q
}

// Level filters by monster level or rank. Supports operator prefixes.
func (q *Query) Level(val string) *Query {
	q.add("level", val)
	return q
}

// Link filters by Link monster rating.
func (q *Query) Link(val int) *Query {
	q.add("link", fmt.Sprintf("%d", val))
	return q
}

// LinkMarker filters by one or more Link arrow markers. Multiple markers are
// AND-ed together.
func (q *Query) LinkMarker(markers ...string) *Query {
	q.add("linkmarker", strings.Join(markers, ","))
	return q
}

// Scale filters by Pendulum monster scale value.
func (q *Query) Scale(val int) *Query {
	q.add("scale", fmt.Sprintf("%d", val))
	return q
}

// CardSet filters cards that appear in the named set (e.g. "Legend of Blue Eyes White Dragon").
func (q *Query) CardSet(set string) *Query {
	q.add("cardset", set)
	return q
}

// Archetype filters cards belonging to the named archetype.
func (q *Query) Archetype(arch string) *Query {
	q.add("archetype", arch)
	return q
}

// BanlistFilter filters by the given banlist.
func (q *Query) BanlistFilter(list enums.Banlist) *Query {
	q.add("banlist", string(list))
	return q
}

// Format filters by game format.
func (q *Query) Format(f enums.Format) *Query {
	q.add("format", string(f))
	return q
}

// Sort sets the field by which results are sorted.
func (q *Query) Sort(field enums.SortOrder) *Query {
	q.add("sort", string(field))
	return q
}

// Misc requests additional miscellaneous card data when yes is true. Has no
// effect when false.
func (q *Query) Misc(yes bool) *Query {
	if yes {
		q.add("misc", "yes")
	}
	return q
}

// Staple filters for cards considered staples when yes is true. Has no effect
// when false.
func (q *Query) Staple(yes bool) *Query {
	if yes {
		q.add("staple", "yes")
	}
	return q
}

// HasEffect filters by whether the card has an effect. Both true and false
// produce an explicit query parameter.
func (q *Query) HasEffect(yes bool) *Query {
	if yes {
		q.add("has_effect", "true")
	} else {
		q.add("has_effect", "false")
	}
	return q
}

// StartDate filters cards released on or after date (format: YYYY-MM-DD or
// as accepted by the API).
func (q *Query) StartDate(date string) *Query {
	q.add("startdate", date)
	return q
}

// EndDate filters cards released on or before date.
func (q *Query) EndDate(date string) *Query {
	q.add("enddate", date)
	return q
}

// DateRegion specifies which date field StartDate/EndDate apply to (e.g.
// "tcg_date").
func (q *Query) DateRegion(region string) *Query {
	q.add("dateregion", region)
	return q
}

// Language requests card data in the given non-English language.
func (q *Query) Language(lang enums.Language) *Query {
	q.add("language", string(lang))
	return q
}

// Num limits the number of results returned (pagination page size).
func (q *Query) Num(n int) *Query {
	q.add("num", fmt.Sprintf("%d", n))
	return q
}

// Offset skips the first n results (pagination offset).
func (q *Query) Offset(n int) *Query {
	q.add("offset", fmt.Sprintf("%d", n))
	return q
}

// Build encodes all accumulated parameters as a URL query string beginning
// with "?". Returns an empty string if no parameters have been set. All keys
// and values are percent-encoded.
func (q *Query) Build() string {
	if len(q.params) == 0 {
		return ""
	}
	var parts []string
	for _, p := range q.params {
		k := url.QueryEscape(p.key)
		v := url.QueryEscape(p.value)
		parts = append(parts, fmt.Sprintf("%s=%s", k, v))
	}
	return "?" + strings.Join(parts, "&")
}
